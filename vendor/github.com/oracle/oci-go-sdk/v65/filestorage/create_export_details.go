// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExportDetails Details for creating the export.
type CreateExportDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's export set.
	ExportSetId *string `mandatory:"true" json:"exportSetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's file system.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// Path used to access the associated file system.
	// Avoid entering confidential information.
	// Example: `/mediafiles`
	Path *string `mandatory:"true" json:"path"`

	// Export options for the new export. If left unspecified,
	// defaults to:
	//        [
	//          {
	//             "source" : "0.0.0.0/0",
	//             "requirePrivilegedSourcePort" : false,
	//             "access" : "READ_WRITE",
	//             "identitySquash" : "NONE"
	//           }
	//        ]
	//   **Note:** Mount targets do not have Internet-routable IP
	//   addresses.  Therefore they will not be reachable from the
	//   Internet, even if an associated `ClientOptions` item has
	//   a source of `0.0.0.0/0`.
	//   **If set to the empty array then the export will not be
	//   visible to any clients.**
	//   The export's `exportOptions` can be changed after creation
	//   using the `UpdateExport` operation.
	ExportOptions []ClientOptions `mandatory:"false" json:"exportOptions"`
}

func (m CreateExportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
