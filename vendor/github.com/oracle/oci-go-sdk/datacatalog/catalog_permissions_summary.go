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

// CatalogPermissionsSummary General permissions object.
type CatalogPermissionsSummary struct {

	// An array of permissions
	UserPermissions []string `mandatory:"false" json:"userPermissions"`

	// The Catalog's Oracle ID (OCID).
	CatalogId *string `mandatory:"false" json:"catalogId"`
}

func (m CatalogPermissionsSummary) String() string {
	return common.PointerString(m)
}
