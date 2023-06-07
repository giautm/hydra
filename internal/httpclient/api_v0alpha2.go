/*
Ory Hydra API

Documentation for all of Ory Hydra's APIs.

API version:
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// V0alpha2ApiService V0alpha2Api service
type V0alpha2ApiService service

type ApiPerformOAuth2DeviceFlowRequest struct {
	ctx        context.Context
	ApiService *V0alpha2ApiService
}

func (r ApiPerformOAuth2DeviceFlowRequest) Execute() (*DeviceAuthorization, *http.Response, error) {
	return r.ApiService.PerformOAuth2DeviceFlowExecute(r)
}

/*
PerformOAuth2DeviceFlow The OAuth 2.0 Device Authorize Endpoint

This endpoint is not documented here because you should never use your own implementation to perform OAuth2 flows.
OAuth2 is a very popular protocol and a library for your programming language will exists.

To learn more about this flow please refer to the specification: https://tools.ietf.org/html/rfc8628

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPerformOAuth2DeviceFlowRequest
*/
func (a *V0alpha2ApiService) PerformOAuth2DeviceFlow(ctx context.Context) ApiPerformOAuth2DeviceFlowRequest {
	return ApiPerformOAuth2DeviceFlowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeviceAuthorization
func (a *V0alpha2ApiService) PerformOAuth2DeviceFlowExecute(r ApiPerformOAuth2DeviceFlowRequest) (*DeviceAuthorization, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeviceAuthorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V0alpha2ApiService.PerformOAuth2DeviceFlow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/device/auth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorOAuth2
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
