// This file was generated by Conjure and should not be manually edited.

package api

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/palantir/conjure-go-runtime/conjure-go-client/httpclient"
	"github.com/palantir/pkg/bearertoken"
	"github.com/palantir/pkg/safelong"
	"github.com/palantir/pkg/uuid"
)

type TestServiceClient interface {
	Echo(ctx context.Context, cookieToken bearertoken.Token) error
	GetPathParam(ctx context.Context, authHeader bearertoken.Token, myPathParamArg string) error
	GetPathParamAlias(ctx context.Context, authHeader bearertoken.Token, myPathParamArg StringAlias) error
	PostPathParam(ctx context.Context, authHeader bearertoken.Token, myPathParam1Arg string, myPathParam2Arg bool, myBodyParamArg CustomObject, myQueryParam1Arg string, myQueryParam2Arg string, myQueryParam3Arg float64, myQueryParam4Arg *safelong.SafeLong, myQueryParam5Arg *string, myHeaderParam1Arg safelong.SafeLong, myHeaderParam2Arg *uuid.UUID) (CustomObject, error)
	Bytes(ctx context.Context) (CustomObject, error)
	GetBinary(ctx context.Context) (io.ReadCloser, error)
	PostBinary(ctx context.Context, myBytesArg func() io.ReadCloser) (io.ReadCloser, error)
	PutBinary(ctx context.Context, myBytesArg func() io.ReadCloser) error
	// An endpoint that uses go keywords
	Chan(ctx context.Context, varArg string, importArg map[string]string, typeArg string, returnArg safelong.SafeLong) error
}

type testServiceClient struct {
	client httpclient.Client
}

func NewTestServiceClient(client httpclient.Client) TestServiceClient {
	return &testServiceClient{client: client}
}

func (c *testServiceClient) Echo(ctx context.Context, cookieToken bearertoken.Token) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Echo"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithHeader("Cookie", fmt.Sprint("PALANTIR_TOKEN=", cookieToken)))
	requestParams = append(requestParams, httpclient.WithPathf("/echo"))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *testServiceClient) GetPathParam(ctx context.Context, authHeader bearertoken.Token, myPathParamArg string) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("GetPathParam"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithHeader("Authorization", fmt.Sprint("Bearer ", authHeader)))
	requestParams = append(requestParams, httpclient.WithPathf("/path/%s", url.PathEscape(fmt.Sprint(myPathParamArg))))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *testServiceClient) GetPathParamAlias(ctx context.Context, authHeader bearertoken.Token, myPathParamArg StringAlias) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("GetPathParamAlias"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithHeader("Authorization", fmt.Sprint("Bearer ", authHeader)))
	requestParams = append(requestParams, httpclient.WithPathf("/path/alias/%s", url.PathEscape(fmt.Sprint(myPathParamArg))))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *testServiceClient) PostPathParam(ctx context.Context, authHeader bearertoken.Token, myPathParam1Arg string, myPathParam2Arg bool, myBodyParamArg CustomObject, myQueryParam1Arg string, myQueryParam2Arg string, myQueryParam3Arg float64, myQueryParam4Arg *safelong.SafeLong, myQueryParam5Arg *string, myHeaderParam1Arg safelong.SafeLong, myHeaderParam2Arg *uuid.UUID) (CustomObject, error) {
	var defaultReturnVal CustomObject
	var returnVal *CustomObject
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("PostPathParam"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("POST"))
	requestParams = append(requestParams, httpclient.WithHeader("Authorization", fmt.Sprint("Bearer ", authHeader)))
	requestParams = append(requestParams, httpclient.WithPathf("/path/%s/%s", url.PathEscape(fmt.Sprint(myPathParam1Arg)), url.PathEscape(fmt.Sprint(myPathParam2Arg))))
	requestParams = append(requestParams, httpclient.WithJSONRequest(myBodyParamArg))
	requestParams = append(requestParams, httpclient.WithHeader("X-My-Header1-Abc", fmt.Sprint(myHeaderParam1Arg)))
	requestParams = append(requestParams, httpclient.WithHeader("X-My-Header2", fmt.Sprint(myHeaderParam2Arg)))
	queryParams := make(url.Values)
	queryParams.Set("query1", fmt.Sprint(myQueryParam1Arg))
	queryParams.Set("myQueryParam2", fmt.Sprint(myQueryParam2Arg))
	queryParams.Set("myQueryParam3", fmt.Sprint(myQueryParam3Arg))
	if myQueryParam4Arg != nil {
		queryParams.Set("myQueryParam4", fmt.Sprint(*myQueryParam4Arg))
	}
	if myQueryParam5Arg != nil {
		queryParams.Set("myQueryParam5", fmt.Sprint(*myQueryParam5Arg))
	}
	requestParams = append(requestParams, httpclient.WithQueryValues(queryParams))
	requestParams = append(requestParams, httpclient.WithJSONResponse(&returnVal))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return defaultReturnVal, err
	}
	_ = resp
	if returnVal == nil {
		return defaultReturnVal, fmt.Errorf("returnVal cannot be nil")
	}
	return *returnVal, nil
}

func (c *testServiceClient) Bytes(ctx context.Context) (CustomObject, error) {
	var defaultReturnVal CustomObject
	var returnVal *CustomObject
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Bytes"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithPathf("/bytes"))
	requestParams = append(requestParams, httpclient.WithJSONResponse(&returnVal))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return defaultReturnVal, err
	}
	_ = resp
	if returnVal == nil {
		return defaultReturnVal, fmt.Errorf("returnVal cannot be nil")
	}
	return *returnVal, nil
}

func (c *testServiceClient) GetBinary(ctx context.Context) (io.ReadCloser, error) {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("GetBinary"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithPathf("/binary"))
	requestParams = append(requestParams, httpclient.WithRawResponseBody())
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *testServiceClient) PostBinary(ctx context.Context, myBytesArg func() io.ReadCloser) (io.ReadCloser, error) {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("PostBinary"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("POST"))
	requestParams = append(requestParams, httpclient.WithPathf("/binary"))
	requestParams = append(requestParams, httpclient.WithRawRequestBodyProvider(myBytesArg))
	requestParams = append(requestParams, httpclient.WithRawResponseBody())
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *testServiceClient) PutBinary(ctx context.Context, myBytesArg func() io.ReadCloser) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("PutBinary"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("PUT"))
	requestParams = append(requestParams, httpclient.WithPathf("/binary"))
	requestParams = append(requestParams, httpclient.WithRawRequestBodyProvider(myBytesArg))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *testServiceClient) Chan(ctx context.Context, varArg string, importArg map[string]string, typeArg string, returnArg safelong.SafeLong) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Chan"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("POST"))
	requestParams = append(requestParams, httpclient.WithPathf("/chan/%s", url.PathEscape(fmt.Sprint(varArg))))
	requestParams = append(requestParams, httpclient.WithJSONRequest(importArg))
	requestParams = append(requestParams, httpclient.WithHeader("X-My-Header2", fmt.Sprint(returnArg)))
	queryParams := make(url.Values)
	queryParams.Set("type", fmt.Sprint(typeArg))
	requestParams = append(requestParams, httpclient.WithQueryValues(queryParams))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

type TestServiceClientWithAuth interface {
	Echo(ctx context.Context) error
	GetPathParam(ctx context.Context, myPathParamArg string) error
	GetPathParamAlias(ctx context.Context, myPathParamArg StringAlias) error
	PostPathParam(ctx context.Context, myPathParam1Arg string, myPathParam2Arg bool, myBodyParamArg CustomObject, myQueryParam1Arg string, myQueryParam2Arg string, myQueryParam3Arg float64, myQueryParam4Arg *safelong.SafeLong, myQueryParam5Arg *string, myHeaderParam1Arg safelong.SafeLong, myHeaderParam2Arg *uuid.UUID) (CustomObject, error)
	Bytes(ctx context.Context) (CustomObject, error)
	GetBinary(ctx context.Context) (io.ReadCloser, error)
	PostBinary(ctx context.Context, myBytesArg func() io.ReadCloser) (io.ReadCloser, error)
	PutBinary(ctx context.Context, myBytesArg func() io.ReadCloser) error
	// An endpoint that uses go keywords
	Chan(ctx context.Context, varArg string, importArg map[string]string, typeArg string, returnArg safelong.SafeLong) error
}

func NewTestServiceClientWithAuth(client TestServiceClient, authHeader bearertoken.Token, cookieToken bearertoken.Token) TestServiceClientWithAuth {
	return &testServiceClientWithAuth{client: client, authHeader: authHeader, cookieToken: cookieToken}
}

type testServiceClientWithAuth struct {
	client      TestServiceClient
	authHeader  bearertoken.Token
	cookieToken bearertoken.Token
}

func (c *testServiceClientWithAuth) Echo(ctx context.Context) error {
	return c.client.Echo(ctx, c.cookieToken)
}

func (c *testServiceClientWithAuth) GetPathParam(ctx context.Context, myPathParamArg string) error {
	return c.client.GetPathParam(ctx, c.authHeader, myPathParamArg)
}

func (c *testServiceClientWithAuth) GetPathParamAlias(ctx context.Context, myPathParamArg StringAlias) error {
	return c.client.GetPathParamAlias(ctx, c.authHeader, myPathParamArg)
}

func (c *testServiceClientWithAuth) PostPathParam(ctx context.Context, myPathParam1Arg string, myPathParam2Arg bool, myBodyParamArg CustomObject, myQueryParam1Arg string, myQueryParam2Arg string, myQueryParam3Arg float64, myQueryParam4Arg *safelong.SafeLong, myQueryParam5Arg *string, myHeaderParam1Arg safelong.SafeLong, myHeaderParam2Arg *uuid.UUID) (CustomObject, error) {
	return c.client.PostPathParam(ctx, c.authHeader, myPathParam1Arg, myPathParam2Arg, myBodyParamArg, myQueryParam1Arg, myQueryParam2Arg, myQueryParam3Arg, myQueryParam4Arg, myQueryParam5Arg, myHeaderParam1Arg, myHeaderParam2Arg)
}

func (c *testServiceClientWithAuth) Bytes(ctx context.Context) (CustomObject, error) {
	return c.client.Bytes(ctx)
}

func (c *testServiceClientWithAuth) GetBinary(ctx context.Context) (io.ReadCloser, error) {
	return c.client.GetBinary(ctx)
}

func (c *testServiceClientWithAuth) PostBinary(ctx context.Context, myBytesArg func() io.ReadCloser) (io.ReadCloser, error) {
	return c.client.PostBinary(ctx, myBytesArg)
}

func (c *testServiceClientWithAuth) PutBinary(ctx context.Context, myBytesArg func() io.ReadCloser) error {
	return c.client.PutBinary(ctx, myBytesArg)
}

func (c *testServiceClientWithAuth) Chan(ctx context.Context, varArg string, importArg map[string]string, typeArg string, returnArg safelong.SafeLong) error {
	return c.client.Chan(ctx, varArg, importArg, typeArg, returnArg)
}
