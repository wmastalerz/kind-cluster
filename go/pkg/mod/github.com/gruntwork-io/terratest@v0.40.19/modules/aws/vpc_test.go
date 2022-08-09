package aws

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestGetDefaultVpc(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	vpc := GetDefaultVpc(t, region)

	assert.NotEmpty(t, vpc.Name)
	assert.True(t, len(vpc.Subnets) > 0)
	assert.Regexp(t, "^vpc-[[:alnum:]]+$", vpc.Id)
}

func TestGetVpcById(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	vpc := createVpc(t, region)
	defer deleteVpc(t, *vpc.VpcId, region)

	vpcTest := GetVpcById(t, *vpc.VpcId, region)
	assert.Equal(t, *vpc.VpcId, vpcTest.Id)
}

func TestGetVpcsE(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	azs := GetAvailabilityZones(t, region)

	isDefaultFilterName := "isDefault"
	isDefaultFilterValue := "true"

	defaultVpcFilter := ec2.Filter{Name: &isDefaultFilterName, Values: []*string{&isDefaultFilterValue}}
	vpcs, _ := GetVpcsE(t, []*ec2.Filter{&defaultVpcFilter}, region)

	require.Equal(t, len(vpcs), 1)
	assert.NotEmpty(t, vpcs[0].Name)

	// the default VPC has by default one subnet per availability zone
	// https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html
	assert.True(t, len(vpcs[0].Subnets) >= len(azs))
}

func TestGetFirstTwoOctets(t *testing.T) {
	t.Parallel()

	firstTwo := GetFirstTwoOctets("10.100.0.0/28")
	if firstTwo != "10.100" {
		t.Errorf("Received: %s, Expected: 10.100", firstTwo)
	}
}

func TestIsPublicSubnet(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	vpc := createVpc(t, region)
	defer deleteVpc(t, *vpc.VpcId, region)

	routeTable := createRouteTable(t, *vpc.VpcId, region)
	subnet := createSubnet(t, *vpc.VpcId, *routeTable.RouteTableId, region)
	assert.False(t, IsPublicSubnet(t, *subnet.SubnetId, region))

	createPublicRoute(t, *vpc.VpcId, *routeTable.RouteTableId, region)
	assert.True(t, IsPublicSubnet(t, *subnet.SubnetId, region))
}

func TestGetDefaultSubnetIDsForVpc(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	defaultVpcBeforeSubnetCreation := GetDefaultVpc(t, region)

	// Creates a subnet in the default VPC with deferred deletion
	// and fetches vpc object again
	subnetName := fmt.Sprintf("%s-subnet", t.Name())
	subnet := createPrivateSubnetInDefaultVpc(t, defaultVpcBeforeSubnetCreation.Id, subnetName, region)
	defer deleteSubnet(t, *subnet.SubnetId, region)
	defaultVpc := GetDefaultVpc(t, region)

	defaultSubnetIDs := GetDefaultSubnetIDsForVpc(t, *defaultVpc)
	assert.NotEmpty(t, defaultSubnetIDs)
	// Checks that the amount of default subnets is smaller than
	// total number of subnets in default vpc
	assert.True(t, len(defaultSubnetIDs) < len(defaultVpc.Subnets))

	availabilityZones := []string{}
	for _, id := range defaultSubnetIDs {
		// check if the recently created subnet does not come up here
		assert.NotEqual(t, id, subnet.SubnetId)
		// default subnets are by default public
		// https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html
		assert.True(t, IsPublicSubnet(t, id, region))
		for _, subnet := range defaultVpc.Subnets {
			if id == subnet.Id {
				availabilityZones = append(availabilityZones, subnet.AvailabilityZone)
			}
		}
	}
	// only one default subnet is allowed per AZ
	uniqueAZs := map[string]bool{}
	for _, az := range availabilityZones {
		uniqueAZs[az] = true
	}
	assert.Equal(t, len(defaultSubnetIDs), len(uniqueAZs))
}

func TestGetTagsForVpc(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	vpc := createVpc(t, region)
	defer deleteVpc(t, *vpc.VpcId, region)

	noTags := GetTagsForVpc(t, *vpc.VpcId, region)
	assert.True(t, len(vpc.Tags) == 0)
	assert.True(t, len(noTags) == 0)

	testTags := make(map[string]string)
	testTags["TagKey1"] = "TagValue1"
	testTags["TagKey2"] = "TagValue2"

	AddTagsToResource(t, region, *vpc.VpcId, testTags)
	vpcWithTags := GetVpcById(t, *vpc.VpcId, region)
	tags := GetTagsForVpc(t, *vpc.VpcId, region)

	assert.True(t, len(vpcWithTags.Tags) == len(testTags))
	assert.True(t, len(tags) == len(testTags))
}

func TestGetTagsForSubnet(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	vpc := createVpc(t, region)
	defer deleteVpc(t, *vpc.VpcId, region)

	routeTable := createRouteTable(t, *vpc.VpcId, region)
	subnet := createSubnet(t, *vpc.VpcId, *routeTable.RouteTableId, region)

	noTags := GetTagsForSubnet(t, *subnet.SubnetId, region)
	assert.True(t, len(subnet.Tags) == 0)
	assert.True(t, len(noTags) == 0)

	testTags := make(map[string]string)
	testTags["TagKey1"] = "TagValue1"
	testTags["TagKey2"] = "TagValue2"

	AddTagsToResource(t, region, *subnet.SubnetId, testTags)

	subnetWithTags := GetSubnetsForVpc(t, *vpc.VpcId, region)[0]
	tags := GetTagsForSubnet(t, *subnet.SubnetId, region)

	assert.True(t, len(subnetWithTags.Tags) == len(testTags))
	assert.True(t, len(tags) == len(testTags))
	assert.True(t, testTags["TagKey1"] == "TagValue1")
	assert.True(t, testTags["TagKey2"] == "TagValue2")
}

func createPublicRoute(t *testing.T, vpcId string, routeTableId string, region string) {
	ec2Client := NewEc2Client(t, region)

	createIGWOut, igerr := ec2Client.CreateInternetGateway(&ec2.CreateInternetGatewayInput{})
	require.NoError(t, igerr)

	_, aigerr := ec2Client.AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		InternetGatewayId: createIGWOut.InternetGateway.InternetGatewayId,
		VpcId:             aws.String(vpcId),
	})
	require.NoError(t, aigerr)

	_, err := ec2Client.CreateRoute(&ec2.CreateRouteInput{
		RouteTableId:         aws.String(routeTableId),
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
		GatewayId:            createIGWOut.InternetGateway.InternetGatewayId,
	})

	require.NoError(t, err)
}

func createRouteTable(t *testing.T, vpcId string, region string) ec2.RouteTable {
	ec2Client := NewEc2Client(t, region)

	createRouteTableOutput, err := ec2Client.CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId: aws.String(vpcId),
	})

	require.NoError(t, err)
	return *createRouteTableOutput.RouteTable
}

func createSubnet(t *testing.T, vpcId string, routeTableId string, region string) ec2.Subnet {
	ec2Client := NewEc2Client(t, region)

	createSubnetOutput, err := ec2Client.CreateSubnet(&ec2.CreateSubnetInput{
		CidrBlock: aws.String("10.10.1.0/24"),
		VpcId:     aws.String(vpcId),
	})
	require.NoError(t, err)

	_, err = ec2Client.AssociateRouteTable(&ec2.AssociateRouteTableInput{
		RouteTableId: aws.String(routeTableId),
		SubnetId:     aws.String(*createSubnetOutput.Subnet.SubnetId),
	})
	require.NoError(t, err)

	return *createSubnetOutput.Subnet
}

func createPrivateSubnetInDefaultVpc(t *testing.T, vpcId string, subnetName string, region string) ec2.Subnet {
	ec2Client := NewEc2Client(t, region)

	createSubnetOutput, err := ec2Client.CreateSubnet(&ec2.CreateSubnetInput{
		CidrBlock: aws.String("172.31.172.0/24"),
		VpcId:     aws.String(vpcId),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("subnet"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String(subnetName),
					},
				},
			},
		},
	})
	require.NoError(t, err)

	return *createSubnetOutput.Subnet
}

func deleteSubnet(t *testing.T, subnetId string, region string) {
	ec2Client := NewEc2Client(t, region)

	_, err := ec2Client.DeleteSubnet(&ec2.DeleteSubnetInput{
		SubnetId: aws.String(subnetId),
	})
	require.NoError(t, err)
}

func createVpc(t *testing.T, region string) ec2.Vpc {
	ec2Client := NewEc2Client(t, region)

	createVpcOutput, err := ec2Client.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: aws.String("10.10.0.0/16"),
	})

	require.NoError(t, err)
	return *createVpcOutput.Vpc
}

func deleteRouteTables(t *testing.T, vpcId string, region string) {
	ec2Client := NewEc2Client(t, region)

	vpcIDFilterName := "vpc-id"
	vpcIDFilter := ec2.Filter{Name: &vpcIDFilterName, Values: []*string{&vpcId}}

	// "You can't delete the main route table."
	mainRTFilterName := "association.main"
	mainRTFilterValue := "false"
	notMainRTFilter := ec2.Filter{Name: &mainRTFilterName, Values: []*string{&mainRTFilterValue}}

	filters := []*ec2.Filter{&vpcIDFilter, &notMainRTFilter}

	rtOutput, err := ec2Client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{Filters: filters})
	require.NoError(t, err)

	for _, rt := range rtOutput.RouteTables {

		// "You must disassociate the route table from any subnets before you can delete it."
		for _, assoc := range rt.Associations {
			_, disassocErr := ec2Client.DisassociateRouteTable(&ec2.DisassociateRouteTableInput{
				AssociationId: assoc.RouteTableAssociationId,
			})
			require.NoError(t, disassocErr)
		}

		_, err := ec2Client.DeleteRouteTable(&ec2.DeleteRouteTableInput{
			RouteTableId: rt.RouteTableId,
		})
		require.NoError(t, err)
	}
}

func deleteSubnets(t *testing.T, vpcId string, region string) {
	ec2Client := NewEc2Client(t, region)
	vpcIDFilterName := "vpc-id"
	vpcIDFilter := ec2.Filter{Name: &vpcIDFilterName, Values: []*string{&vpcId}}

	subnetsOutput, err := ec2Client.DescribeSubnets(&ec2.DescribeSubnetsInput{Filters: []*ec2.Filter{&vpcIDFilter}})
	require.NoError(t, err)

	for _, subnet := range subnetsOutput.Subnets {
		_, err := ec2Client.DeleteSubnet(&ec2.DeleteSubnetInput{
			SubnetId: subnet.SubnetId,
		})
		require.NoError(t, err)
	}
}

func deleteInternetGateways(t *testing.T, vpcId string, region string) {
	ec2Client := NewEc2Client(t, region)
	vpcIDFilterName := "attachment.vpc-id"
	vpcIDFilter := ec2.Filter{Name: &vpcIDFilterName, Values: []*string{&vpcId}}

	igwOutput, err := ec2Client.DescribeInternetGateways(&ec2.DescribeInternetGatewaysInput{Filters: []*ec2.Filter{&vpcIDFilter}})
	require.NoError(t, err)

	for _, igw := range igwOutput.InternetGateways {

		_, detachErr := ec2Client.DetachInternetGateway(&ec2.DetachInternetGatewayInput{
			InternetGatewayId: igw.InternetGatewayId,
			VpcId:             aws.String(vpcId),
		})
		require.NoError(t, detachErr)

		_, err := ec2Client.DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
			InternetGatewayId: igw.InternetGatewayId,
		})
		require.NoError(t, err)
	}
}

func deleteVpc(t *testing.T, vpcId string, region string) {
	ec2Client := NewEc2Client(t, region)

	deleteRouteTables(t, vpcId, region)
	deleteSubnets(t, vpcId, region)
	deleteInternetGateways(t, vpcId, region)

	_, err := ec2Client.DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: aws.String(vpcId),
	})
	require.NoError(t, err)
}
