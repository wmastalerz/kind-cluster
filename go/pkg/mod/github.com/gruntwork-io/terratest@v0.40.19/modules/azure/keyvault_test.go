package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete key vault resources are added, these tests can be extended.
*/

func TestKeyVaultSecretExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultSecretName := "fakeSecretName"
	_, err := KeyVaultSecretExistsE(testKeyVaultName, testKeyVaultSecretName)
	require.Error(t, err)
}

func TestKeyVaultKeyExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultKeyName := "fakeKeyName"
	_, err := KeyVaultKeyExistsE(testKeyVaultName, testKeyVaultKeyName)
	require.Error(t, err)
}

func TestKeyVaultCertificateExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultCertName := "fakeCertName"
	_, err := KeyVaultCertificateExistsE(testKeyVaultName, testKeyVaultCertName)
	require.Error(t, err)
}

func TestGetKeyVault(t *testing.T) {
	t.Parallel()

	resGroupName := ""
	keyVaultName := ""
	subscriptionID := ""

	_, err := GetKeyVaultE(t, resGroupName, keyVaultName, subscriptionID)
	require.Error(t, err)
}
