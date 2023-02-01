// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For information about monitoring, see Monitoring Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlarmSummary A summary of properties for the specified alarm.
// For information about alarms, see Alarms Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#AlarmsOverview).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/sdks.htm).
type AlarmSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the alarm. It does not have to be unique, and it's changeable.
	// This name is sent as the title for notifications related to this alarm.
	// Example: `High CPU Utilization`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric
	// being evaluated by the alarm.
	MetricCompartmentId *string `mandatory:"true" json:"metricCompartmentId"`

	// The source service or application emitting the metric that is evaluated by the alarm.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of
	// the Monitoring service interprets results for each returned time series as Boolean values,
	// where zero represents false and a non-zero value represents true. A true value means that the trigger
	// rule condition has been met. The query must specify a metric, statistic, interval, and trigger
	// rule (threshold or absence). Supported values for interval depend on the specified time range. More
	// interval values are supported for smaller time ranges. Supported grouping functions: `grouping()`, `groupBy()`.
	// For details about Monitoring Query Language (MQL), see Monitoring Query Language (MQL) Reference (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm).
	// For available dimensions, review the metric definition for the supported service.
	// See Supported Services (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).
	// Example of threshold alarm:
	//   -----
	//     CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85
	//   -----
	// Example of absence alarm:
	//   -----
	//     CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()
	//   -----
	Query *string `mandatory:"true" json:"query"`

	// The perceived severity of the alarm with regard to the affected system.
	// Example: `CRITICAL`
	Severity AlarmSummarySeverityEnum `mandatory:"true" json:"severity"`

	// A list of destinations to which the notifications for this alarm will be delivered.
	// Each destination is represented by an OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) related to the supported destination service.
	// For example, a destination using the Notifications service is represented by a topic OCID.
	// Supported destination services: Notifications Service. Limit: One destination per supported destination service.
	Destinations []string `mandatory:"true" json:"destinations"`

	// Whether the alarm is enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The current lifecycle state of the alarm.
	// Example: `DELETED`
	LifecycleState AlarmLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The configuration details for suppressing an alarm.
	Suppression *Suppression `mandatory:"false" json:"suppression"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AlarmSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmSummarySeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlarmSummarySeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlarmLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlarmLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlarmSummarySeverityEnum Enum with underlying type: string
type AlarmSummarySeverityEnum string

// Set of constants representing the allowable values for AlarmSummarySeverityEnum
const (
	AlarmSummarySeverityCritical AlarmSummarySeverityEnum = "CRITICAL"
	AlarmSummarySeverityError    AlarmSummarySeverityEnum = "ERROR"
	AlarmSummarySeverityWarning  AlarmSummarySeverityEnum = "WARNING"
	AlarmSummarySeverityInfo     AlarmSummarySeverityEnum = "INFO"
)

var mappingAlarmSummarySeverityEnum = map[string]AlarmSummarySeverityEnum{
	"CRITICAL": AlarmSummarySeverityCritical,
	"ERROR":    AlarmSummarySeverityError,
	"WARNING":  AlarmSummarySeverityWarning,
	"INFO":     AlarmSummarySeverityInfo,
}

var mappingAlarmSummarySeverityEnumLowerCase = map[string]AlarmSummarySeverityEnum{
	"critical": AlarmSummarySeverityCritical,
	"error":    AlarmSummarySeverityError,
	"warning":  AlarmSummarySeverityWarning,
	"info":     AlarmSummarySeverityInfo,
}

// GetAlarmSummarySeverityEnumValues Enumerates the set of values for AlarmSummarySeverityEnum
func GetAlarmSummarySeverityEnumValues() []AlarmSummarySeverityEnum {
	values := make([]AlarmSummarySeverityEnum, 0)
	for _, v := range mappingAlarmSummarySeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmSummarySeverityEnumStringValues Enumerates the set of values in String for AlarmSummarySeverityEnum
func GetAlarmSummarySeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"ERROR",
		"WARNING",
		"INFO",
	}
}

// GetMappingAlarmSummarySeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmSummarySeverityEnum(val string) (AlarmSummarySeverityEnum, bool) {
	enum, ok := mappingAlarmSummarySeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
