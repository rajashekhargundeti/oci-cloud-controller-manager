// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"fmt"
	"net"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/pkg/errors"
)

// NetworkingInterface defines the subset of the OCI compute API utilised by the CCM.
type NetworkingInterface interface {
	GetSubnet(ctx context.Context, id string) (*core.Subnet, error)
	GetSubnetFromCacheByIP(ip string) (*core.Subnet, error)

	GetSecurityList(ctx context.Context, id string) (core.GetSecurityListResponse, error)
	UpdateSecurityList(ctx context.Context, request core.UpdateSecurityListRequest) (core.UpdateSecurityListResponse, error)

	GetRouteTable(ctx context.Context, routeID string) (core.GetRouteTableResponse, error)
	UpdateRouteTable(ctx context.Context, routeID string, routeRules []core.RouteRule) error

	GetIPFromOCID(ctx context.Context, ipID string) (core.GetPrivateIpResponse, error)
	GetOCIDFromIP(ctx context.Context, ipAddress string, subnetID string) (string, error)

	UpdateVnic(ctx context.Context, vnicID string, skipSourceCheck bool) error
}

func (c *client) getVNIC(ctx context.Context, id string) (*core.Vnic, error) {
	resp, err := c.network.GetVnic(ctx, core.GetVnicRequest{
		VnicId: &id,
	})

	incRequestCounter(err, getVerb, vnicResource)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &resp.Vnic, nil
}

func (c *client) GetSubnet(ctx context.Context, id string) (*core.Subnet, error) {
	item, exists, err := c.subnetCache.GetByKey(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if exists {
		return item.(*core.Subnet), nil
	}

	resp, err := c.network.GetSubnet(ctx, core.GetSubnetRequest{
		SubnetId: &id,
	})
	incRequestCounter(err, getVerb, subnetResource)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	subnet := &resp.Subnet
	_ = c.subnetCache.Add(subnet)
	return subnet, nil
}

// GetSubnetFromCacheByIP checks to see if the given IP is contained by any subnet CIDR block in the subnet cache
// If no hits were found then no subnet and no error will be returned (nil, nil)
func (c *client) GetSubnetFromCacheByIP(ip string) (*core.Subnet, error) {
	ipAddr := net.ParseIP(ip)
	for _, subnetItem := range c.subnetCache.List() {
		subnet := subnetItem.(*core.Subnet)
		_, cidr, err := net.ParseCIDR(*subnet.CidrBlock)
		if err != nil {
			// This should never actually error but just in case
			return nil, fmt.Errorf("unable to parse CIDR block %q for subnet %q: %v", *subnet.CidrBlock, *subnet.Id, err)
		}

		if cidr.Contains(ipAddr) {
			return subnet, nil
		}
	}
	return nil, nil
}

func (c *client) GetSecurityList(ctx context.Context, id string) (core.GetSecurityListResponse, error) {
	resp, err := c.network.GetSecurityList(ctx, core.GetSecurityListRequest{
		SecurityListId: &id,
	})
	incRequestCounter(err, getVerb, securityListResource)

	return resp, errors.WithStack(err)
}

func (c *client) UpdateSecurityList(ctx context.Context, request core.UpdateSecurityListRequest) (core.UpdateSecurityListResponse, error) {
	resp, err := c.network.UpdateSecurityList(ctx, request)
	incRequestCounter(err, updateVerb, securityListResource)
	return resp, errors.WithStack(err)
}

func subnetCacheKeyFn(obj interface{}) (string, error) {
	return *obj.(*core.Subnet).Id, nil
}

func (c *client) GetRouteTable(ctx context.Context, routeID string) (core.GetRouteTableResponse, error) {
	resp, err := c.network.GetRouteTable(ctx, core.GetRouteTableRequest{
		RtId: &routeID,
	})
	incRequestCounter(err, getVerb, routeTableResource)

	if err != nil {
		return core.GetRouteTableResponse{}, err
	}

	return resp, nil
}

func (c *client) GetIPFromOCID(ctx context.Context, ipID string) (core.GetPrivateIpResponse, error) {
	resp, err := c.network.GetPrivateIp(ctx, core.GetPrivateIpRequest{
		PrivateIpId: &ipID,
	})
	incRequestCounter(err, getVerb, privateIPResource)

	if err != nil {
		return core.GetPrivateIpResponse{}, err
	}

	return resp, nil
}

func (c *client) GetOCIDFromIP(ctx context.Context, ipAddress string, subnetID string) (string, error) {
	resp, err := c.network.ListPrivateIps(ctx, core.ListPrivateIpsRequest{
		IpAddress: &ipAddress,
		SubnetId:  &subnetID,
	})
	incRequestCounter(err, getVerb, privateIPResource)

	if err != nil {
		return "", err
	}

	for _, entry := range resp.Items {
		if *entry.IpAddress == ipAddress {
			return *entry.Id, nil
		}
	}

	return "", errors.New("IP not found in subnet")
}

func (c *client) UpdateRouteTable(ctx context.Context, routeID string, routeRules []core.RouteRule) error {
	_, err := c.network.UpdateRouteTable(ctx, core.UpdateRouteTableRequest{
		RtId: &routeID,
		UpdateRouteTableDetails: core.UpdateRouteTableDetails{
			RouteRules: routeRules,
		},
	})
	incRequestCounter(err, updateVerb, routeTableResource)

	return err
}

func (c *client) UpdateVnic(ctx context.Context, vnicID string, skipSourceCheck bool) error {
	_, err := c.network.UpdateVnic(ctx, core.UpdateVnicRequest{
		VnicId: &vnicID,
		UpdateVnicDetails: core.UpdateVnicDetails{
			SkipSourceDestCheck: &skipSourceCheck,
		},
	})
	return err
}
