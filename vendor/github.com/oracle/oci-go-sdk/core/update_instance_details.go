// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateInstanceDetails The representation of UpdateInstanceDetails
type UpdateInstanceDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My bare metal instance`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Instance agent configuration options to choose for updating the instance
	AgentConfig *UpdateInstanceAgentConfigDetails `mandatory:"false" json:"agentConfig"`

	// Custom metadata key/value string pairs that you provide. Any set of key/value pairs
	// provided here will completely replace the current set of key/value pairs in the 'metadata'
	// field on the instance.
	// Both the 'user_data' and 'ssh_authorized_keys' fields cannot be changed after an instance
	// has launched. Any request which updates, removes, or adds either of these fields will be
	// rejected. You must provide the same values for 'user_data' and 'ssh_authorized_keys' that
	// already exist on the instance.
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// Additional metadata key/value pairs that you provide. They serve the same purpose and
	// functionality as fields in the 'metadata' object.
	// They are distinguished from 'metadata' fields in that these can be nested JSON objects
	// (whereas 'metadata' fields are string/string maps only).
	// Both the 'user_data' and 'ssh_authorized_keys' fields cannot be changed after an instance
	// has launched. Any request which updates, removes, or adds either of these fields will be
	// rejected. You must provide the same values for 'user_data' and 'ssh_authorized_keys' that
	// already exist on the instance.
	ExtendedMetadata map[string]interface{} `mandatory:"false" json:"extendedMetadata"`

	// The shape of the instance. The shape determines the number of CPUs and the amount of memory
	// allocated to the instance. You can enumerate all available shapes by calling
	// ListShapes.
	// Example: `VM.Standard1.1`
	Shape *string `mandatory:"false" json:"shape"`

	ShapeConfig *UpdateInstanceShapeConfigDetails `mandatory:"false" json:"shapeConfig"`

	// The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
	// * `LIVE_MIGRATE` - Run maintenance using a live migration.
	// * `REBOOT` - Run maintenance using a reboot.
	PreferredMaintenanceAction UpdateInstanceDetailsPreferredMaintenanceActionEnum `mandatory:"false" json:"preferredMaintenanceAction,omitempty"`
}

func (m UpdateInstanceDetails) String() string {
	return common.PointerString(m)
}

// UpdateInstanceDetailsPreferredMaintenanceActionEnum Enum with underlying type: string
type UpdateInstanceDetailsPreferredMaintenanceActionEnum string

// Set of constants representing the allowable values for UpdateInstanceDetailsPreferredMaintenanceActionEnum
const (
	UpdateInstanceDetailsPreferredMaintenanceActionLiveMigrate UpdateInstanceDetailsPreferredMaintenanceActionEnum = "LIVE_MIGRATE"
	UpdateInstanceDetailsPreferredMaintenanceActionReboot      UpdateInstanceDetailsPreferredMaintenanceActionEnum = "REBOOT"
)

var mappingUpdateInstanceDetailsPreferredMaintenanceAction = map[string]UpdateInstanceDetailsPreferredMaintenanceActionEnum{
	"LIVE_MIGRATE": UpdateInstanceDetailsPreferredMaintenanceActionLiveMigrate,
	"REBOOT":       UpdateInstanceDetailsPreferredMaintenanceActionReboot,
}

// GetUpdateInstanceDetailsPreferredMaintenanceActionEnumValues Enumerates the set of values for UpdateInstanceDetailsPreferredMaintenanceActionEnum
func GetUpdateInstanceDetailsPreferredMaintenanceActionEnumValues() []UpdateInstanceDetailsPreferredMaintenanceActionEnum {
	values := make([]UpdateInstanceDetailsPreferredMaintenanceActionEnum, 0)
	for _, v := range mappingUpdateInstanceDetailsPreferredMaintenanceAction {
		values = append(values, v)
	}
	return values
}
