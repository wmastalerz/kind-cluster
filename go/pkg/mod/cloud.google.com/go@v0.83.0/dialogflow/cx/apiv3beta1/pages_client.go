// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newPagesClientHook clientHook

// PagesCallOptions contains the retry settings for each method of PagesClient.
type PagesCallOptions struct {
	ListPages  []gax.CallOption
	GetPage    []gax.CallOption
	CreatePage []gax.CallOption
	UpdatePage []gax.CallOption
	DeletePage []gax.CallOption
}

func defaultPagesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPagesCallOptions() *PagesCallOptions {
	return &PagesCallOptions{
		ListPages: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetPage: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreatePage: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdatePage: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeletePage: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalPagesClient is an interface that defines the methods availaible from Dialogflow API.
type internalPagesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListPages(context.Context, *cxpb.ListPagesRequest, ...gax.CallOption) *PageIterator
	GetPage(context.Context, *cxpb.GetPageRequest, ...gax.CallOption) (*cxpb.Page, error)
	CreatePage(context.Context, *cxpb.CreatePageRequest, ...gax.CallOption) (*cxpb.Page, error)
	UpdatePage(context.Context, *cxpb.UpdatePageRequest, ...gax.CallOption) (*cxpb.Page, error)
	DeletePage(context.Context, *cxpb.DeletePageRequest, ...gax.CallOption) error
}

// PagesClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing [Pages][google.cloud.dialogflow.cx.v3beta1.Page (at http://google.cloud.dialogflow.cx.v3beta1.Page)].
type PagesClient struct {
	// The internal transport-dependent client.
	internalClient internalPagesClient

	// The call options for this service.
	CallOptions *PagesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PagesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PagesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PagesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListPages returns the list of all pages in the specified flow.
func (c *PagesClient) ListPages(ctx context.Context, req *cxpb.ListPagesRequest, opts ...gax.CallOption) *PageIterator {
	return c.internalClient.ListPages(ctx, req, opts...)
}

// GetPage retrieves the specified page.
func (c *PagesClient) GetPage(ctx context.Context, req *cxpb.GetPageRequest, opts ...gax.CallOption) (*cxpb.Page, error) {
	return c.internalClient.GetPage(ctx, req, opts...)
}

// CreatePage creates a page in the specified flow.
func (c *PagesClient) CreatePage(ctx context.Context, req *cxpb.CreatePageRequest, opts ...gax.CallOption) (*cxpb.Page, error) {
	return c.internalClient.CreatePage(ctx, req, opts...)
}

// UpdatePage updates the specified page.
func (c *PagesClient) UpdatePage(ctx context.Context, req *cxpb.UpdatePageRequest, opts ...gax.CallOption) (*cxpb.Page, error) {
	return c.internalClient.UpdatePage(ctx, req, opts...)
}

// DeletePage deletes the specified page.
func (c *PagesClient) DeletePage(ctx context.Context, req *cxpb.DeletePageRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePage(ctx, req, opts...)
}

// pagesGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type pagesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PagesClient
	CallOptions **PagesCallOptions

	// The gRPC API client.
	pagesClient cxpb.PagesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPagesClient creates a new pages client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing [Pages][google.cloud.dialogflow.cx.v3beta1.Page (at http://google.cloud.dialogflow.cx.v3beta1.Page)].
func NewPagesClient(ctx context.Context, opts ...option.ClientOption) (*PagesClient, error) {
	clientOpts := defaultPagesGRPCClientOptions()
	if newPagesClientHook != nil {
		hookOpts, err := newPagesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PagesClient{CallOptions: defaultPagesCallOptions()}

	c := &pagesGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		pagesClient:      cxpb.NewPagesClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *pagesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *pagesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *pagesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *pagesGRPCClient) ListPages(ctx context.Context, req *cxpb.ListPagesRequest, opts ...gax.CallOption) *PageIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListPages[0:len((*c.CallOptions).ListPages):len((*c.CallOptions).ListPages)], opts...)
	it := &PageIterator{}
	req = proto.Clone(req).(*cxpb.ListPagesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.Page, string, error) {
		var resp *cxpb.ListPagesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.pagesClient.ListPages(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPages(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

func (c *pagesGRPCClient) GetPage(ctx context.Context, req *cxpb.GetPageRequest, opts ...gax.CallOption) (*cxpb.Page, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetPage[0:len((*c.CallOptions).GetPage):len((*c.CallOptions).GetPage)], opts...)
	var resp *cxpb.Page
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.pagesClient.GetPage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *pagesGRPCClient) CreatePage(ctx context.Context, req *cxpb.CreatePageRequest, opts ...gax.CallOption) (*cxpb.Page, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreatePage[0:len((*c.CallOptions).CreatePage):len((*c.CallOptions).CreatePage)], opts...)
	var resp *cxpb.Page
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.pagesClient.CreatePage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *pagesGRPCClient) UpdatePage(ctx context.Context, req *cxpb.UpdatePageRequest, opts ...gax.CallOption) (*cxpb.Page, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "page.name", url.QueryEscape(req.GetPage().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdatePage[0:len((*c.CallOptions).UpdatePage):len((*c.CallOptions).UpdatePage)], opts...)
	var resp *cxpb.Page
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.pagesClient.UpdatePage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *pagesGRPCClient) DeletePage(ctx context.Context, req *cxpb.DeletePageRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeletePage[0:len((*c.CallOptions).DeletePage):len((*c.CallOptions).DeletePage)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.pagesClient.DeletePage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// PageIterator manages a stream of *cxpb.Page.
type PageIterator struct {
	items    []*cxpb.Page
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.Page, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PageIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PageIterator) Next() (*cxpb.Page, error) {
	var item *cxpb.Page
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PageIterator) bufLen() int {
	return len(it.items)
}

func (it *PageIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
