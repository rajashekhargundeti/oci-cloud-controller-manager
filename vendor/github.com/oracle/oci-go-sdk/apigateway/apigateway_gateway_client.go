// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
//

package apigateway

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//GatewayClient a client for Gateway
type GatewayClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGatewayClientWithConfigurationProvider Creates a new default Gateway client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGatewayClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GatewayClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = GatewayClient{BaseClient: baseClient}
	client.BasePath = "20190501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GatewayClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apigateway", "https://apigateway.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GatewayClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *GatewayClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeGatewayCompartment Change the gateway compartment.
func (client GatewayClient) ChangeGatewayCompartment(ctx context.Context, request ChangeGatewayCompartmentRequest) (response ChangeGatewayCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeGatewayCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			response = ChangeGatewayCompartmentResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeGatewayCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeGatewayCompartmentResponse")
	}
	return
}

// changeGatewayCompartment implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) changeGatewayCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gateways/{gatewayId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeGatewayCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGateway Creates a new Gateway.
func (client GatewayClient) CreateGateway(ctx context.Context, request CreateGatewayRequest) (response CreateGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createGateway, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateGatewayResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGatewayResponse")
	}
	return
}

// createGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) createGateway(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gateways")
	if err != nil {
		return nil, err
	}

	var response CreateGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGateway Deletes a Gateway resource by identifier
func (client GatewayClient) DeleteGateway(ctx context.Context, request DeleteGatewayRequest) (response DeleteGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGateway, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteGatewayResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGatewayResponse")
	}
	return
}

// deleteGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) deleteGateway(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/gateways/{gatewayId}")
	if err != nil {
		return nil, err
	}

	var response DeleteGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGateway Gets a Gateway by identifier
func (client GatewayClient) GetGateway(ctx context.Context, request GetGatewayRequest) (response GetGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGateway, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetGatewayResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGatewayResponse")
	}
	return
}

// getGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) getGateway(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gateways/{gatewayId}")
	if err != nil {
		return nil, err
	}

	var response GetGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGateways Returns a list of Gateways.
func (client GatewayClient) ListGateways(ctx context.Context, request ListGatewaysRequest) (response ListGatewaysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGateways, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListGatewaysResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGatewaysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGatewaysResponse")
	}
	return
}

// listGateways implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) listGateways(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gateways")
	if err != nil {
		return nil, err
	}

	var response ListGatewaysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGateway Updates the Gateway
func (client GatewayClient) UpdateGateway(ctx context.Context, request UpdateGatewayRequest) (response UpdateGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGateway, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateGatewayResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGatewayResponse")
	}
	return
}

// updateGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) updateGateway(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/gateways/{gatewayId}")
	if err != nil {
		return nil, err
	}

	var response UpdateGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
