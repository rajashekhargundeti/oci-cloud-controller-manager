// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetCloudSyncRequest wrapper for the GetCloudSync operation
type GetCloudSyncRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the storage gateway.
	StorageGatewayId *string `mandatory:"true" contributesTo:"path" name:"storageGatewayId"`

	// The cloud sync's unique name.
	// Example: `cloud_sync_52019`
	CloudSyncName *string `mandatory:"true" contributesTo:"path" name:"cloudSyncName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCloudSyncRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCloudSyncRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCloudSyncRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetCloudSyncResponse wrapper for the GetCloudSync operation
type GetCloudSyncResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CloudSync instance
	CloudSync `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCloudSyncResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCloudSyncResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
