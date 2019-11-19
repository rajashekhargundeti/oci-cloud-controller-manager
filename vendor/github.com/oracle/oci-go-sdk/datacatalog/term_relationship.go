// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TermRelationship Full Term Relationship Definition. Business term relationship between two terms in a business glossary.
type TermRelationship struct {

	// Unique Term Relationship key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.This is the same as relationshipType for termRelationship
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the Term Relationship usually defined at the time of creation.
	Description *string `mandatory:"false" json:"description"`

	// Unique id of the related term.
	RelatedTermKey *string `mandatory:"false" json:"relatedTermKey"`

	// Name of the related term.
	RelatedTermDisplayName *string `mandatory:"false" json:"relatedTermDisplayName"`

	// Description of the related term.
	RelatedTermDescription *string `mandatory:"false" json:"relatedTermDescription"`

	// URI to the Term Relationship instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// This relationships parent term key.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Name of the parent term.
	ParentTermDisplayName *string `mandatory:"false" json:"parentTermDisplayName"`

	// Description of the parent term.
	ParentTermDescription *string `mandatory:"false" json:"parentTermDescription"`

	// The date and time the Term Relationship was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// State of the Term Relationship.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TermRelationship) String() string {
	return common.PointerString(m)
}
