package costmanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2019-10-01/costmanagement"

// ExecutionStatus enumerates the values for execution status.
type ExecutionStatus string

const (
	// Completed ...
	Completed ExecutionStatus = "Completed"
	// DataNotAvailable ...
	DataNotAvailable ExecutionStatus = "DataNotAvailable"
	// Failed ...
	Failed ExecutionStatus = "Failed"
	// InProgress ...
	InProgress ExecutionStatus = "InProgress"
	// NewDataNotAvailable ...
	NewDataNotAvailable ExecutionStatus = "NewDataNotAvailable"
	// Queued ...
	Queued ExecutionStatus = "Queued"
	// Timeout ...
	Timeout ExecutionStatus = "Timeout"
)

// PossibleExecutionStatusValues returns an array of possible values for the ExecutionStatus const type.
func PossibleExecutionStatusValues() []ExecutionStatus {
	return []ExecutionStatus{Completed, DataNotAvailable, Failed, InProgress, NewDataNotAvailable, Queued, Timeout}
}

// ExecutionType enumerates the values for execution type.
type ExecutionType string

const (
	// OnDemand ...
	OnDemand ExecutionType = "OnDemand"
	// Scheduled ...
	Scheduled ExecutionType = "Scheduled"
)

// PossibleExecutionTypeValues returns an array of possible values for the ExecutionType const type.
func PossibleExecutionTypeValues() []ExecutionType {
	return []ExecutionType{OnDemand, Scheduled}
}

// FormatType enumerates the values for format type.
type FormatType string

const (
	// Csv ...
	Csv FormatType = "Csv"
)

// PossibleFormatTypeValues returns an array of possible values for the FormatType const type.
func PossibleFormatTypeValues() []FormatType {
	return []FormatType{Csv}
}

// GranularityType enumerates the values for granularity type.
type GranularityType string

const (
	// Daily ...
	Daily GranularityType = "Daily"
	// Hourly ...
	Hourly GranularityType = "Hourly"
)

// PossibleGranularityTypeValues returns an array of possible values for the GranularityType const type.
func PossibleGranularityTypeValues() []GranularityType {
	return []GranularityType{Daily, Hourly}
}

// QueryColumnType enumerates the values for query column type.
type QueryColumnType string

const (
	// QueryColumnTypeDimension ...
	QueryColumnTypeDimension QueryColumnType = "Dimension"
	// QueryColumnTypeTag ...
	QueryColumnTypeTag QueryColumnType = "Tag"
)

// PossibleQueryColumnTypeValues returns an array of possible values for the QueryColumnType const type.
func PossibleQueryColumnTypeValues() []QueryColumnType {
	return []QueryColumnType{QueryColumnTypeDimension, QueryColumnTypeTag}
}

// RecurrenceType enumerates the values for recurrence type.
type RecurrenceType string

const (
	// RecurrenceTypeAnnually ...
	RecurrenceTypeAnnually RecurrenceType = "Annually"
	// RecurrenceTypeDaily ...
	RecurrenceTypeDaily RecurrenceType = "Daily"
	// RecurrenceTypeMonthly ...
	RecurrenceTypeMonthly RecurrenceType = "Monthly"
	// RecurrenceTypeWeekly ...
	RecurrenceTypeWeekly RecurrenceType = "Weekly"
)

// PossibleRecurrenceTypeValues returns an array of possible values for the RecurrenceType const type.
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return []RecurrenceType{RecurrenceTypeAnnually, RecurrenceTypeDaily, RecurrenceTypeMonthly, RecurrenceTypeWeekly}
}

// SortDirection enumerates the values for sort direction.
type SortDirection string

const (
	// Ascending ...
	Ascending SortDirection = "Ascending"
	// Descending ...
	Descending SortDirection = "Descending"
)

// PossibleSortDirectionValues returns an array of possible values for the SortDirection const type.
func PossibleSortDirectionValues() []SortDirection {
	return []SortDirection{Ascending, Descending}
}

// StatusType enumerates the values for status type.
type StatusType string

const (
	// Active ...
	Active StatusType = "Active"
	// Inactive ...
	Inactive StatusType = "Inactive"
)

// PossibleStatusTypeValues returns an array of possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return []StatusType{Active, Inactive}
}

// TimeframeType enumerates the values for timeframe type.
type TimeframeType string

const (
	// Custom ...
	Custom TimeframeType = "Custom"
	// MonthToDate ...
	MonthToDate TimeframeType = "MonthToDate"
	// TheLastMonth ...
	TheLastMonth TimeframeType = "TheLastMonth"
	// TheLastWeek ...
	TheLastWeek TimeframeType = "TheLastWeek"
	// TheLastYear ...
	TheLastYear TimeframeType = "TheLastYear"
	// WeekToDate ...
	WeekToDate TimeframeType = "WeekToDate"
	// YearToDate ...
	YearToDate TimeframeType = "YearToDate"
)

// PossibleTimeframeTypeValues returns an array of possible values for the TimeframeType const type.
func PossibleTimeframeTypeValues() []TimeframeType {
	return []TimeframeType{Custom, MonthToDate, TheLastMonth, TheLastWeek, TheLastYear, WeekToDate, YearToDate}
}

// CommonExportProperties the common properties of the export.
type CommonExportProperties struct {
	// Format - The format of the export being delivered. Possible values include: 'Csv'
	Format FormatType `json:"format,omitempty"`
	// DeliveryInfo - Has delivery information for the export.
	DeliveryInfo *ExportDeliveryInfo `json:"deliveryInfo,omitempty"`
	// Definition - Has definition for the export.
	Definition *QueryDefinition `json:"definition,omitempty"`
}

// Dimension ...
type Dimension struct {
	*DimensionProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Dimension.
func (d Dimension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DimensionProperties != nil {
		objectMap["properties"] = d.DimensionProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Dimension struct.
func (d *Dimension) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var dimensionProperties DimensionProperties
				err = json.Unmarshal(*v, &dimensionProperties)
				if err != nil {
					return err
				}
				d.DimensionProperties = &dimensionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				d.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				d.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				d.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				d.Tags = tags
			}
		}
	}

	return nil
}

// DimensionProperties ...
type DimensionProperties struct {
	// Description - READ-ONLY; Dimension description.
	Description *string `json:"description,omitempty"`
	// FilterEnabled - READ-ONLY; Filter enabled.
	FilterEnabled *bool `json:"filterEnabled,omitempty"`
	// GroupingEnabled - READ-ONLY; Grouping enabled.
	GroupingEnabled *bool     `json:"groupingEnabled,omitempty"`
	Data            *[]string `json:"data,omitempty"`
	// Total - READ-ONLY; Total number of data for the dimension.
	Total *int32 `json:"total,omitempty"`
	// Category - READ-ONLY; Dimension category.
	Category *string `json:"category,omitempty"`
	// UsageStart - READ-ONLY; Usage start.
	UsageStart *date.Time `json:"usageStart,omitempty"`
	// UsageEnd - READ-ONLY; Usage end.
	UsageEnd *date.Time `json:"usageEnd,omitempty"`
	// NextLink - READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// DimensionsListResult result of listing dimensions. It contains a list of available dimensions.
type DimensionsListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of dimensions.
	Value *[]Dimension `json:"value,omitempty"`
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - READ-ONLY; Error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The
// reason is provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// Export a export resource.
type Export struct {
	autorest.Response `json:"-"`
	*ExportProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Export.
func (e Export) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if e.ExportProperties != nil {
		objectMap["properties"] = e.ExportProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Export struct.
func (e *Export) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var exportProperties ExportProperties
				err = json.Unmarshal(*v, &exportProperties)
				if err != nil {
					return err
				}
				e.ExportProperties = &exportProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				e.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				e.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				e.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				e.Tags = tags
			}
		}
	}

	return nil
}

// ExportDeliveryDestination the destination information for the delivery of the export.
type ExportDeliveryDestination struct {
	// ResourceID - The resource id of the storage account where exports will be delivered.
	ResourceID *string `json:"resourceId,omitempty"`
	// Container - The name of the container where exports will be uploaded.
	Container *string `json:"container,omitempty"`
	// RootFolderPath - The name of the directory where exports will be uploaded.
	RootFolderPath *string `json:"rootFolderPath,omitempty"`
}

// ExportDeliveryInfo the delivery information associated with a export.
type ExportDeliveryInfo struct {
	// Destination - Has destination for the export being delivered.
	Destination *ExportDeliveryDestination `json:"destination,omitempty"`
}

// ExportExecution a export execution.
type ExportExecution struct {
	*ExportExecutionProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ExportExecution.
func (ee ExportExecution) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ee.ExportExecutionProperties != nil {
		objectMap["properties"] = ee.ExportExecutionProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ExportExecution struct.
func (ee *ExportExecution) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var exportExecutionProperties ExportExecutionProperties
				err = json.Unmarshal(*v, &exportExecutionProperties)
				if err != nil {
					return err
				}
				ee.ExportExecutionProperties = &exportExecutionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ee.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ee.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ee.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ee.Tags = tags
			}
		}
	}

	return nil
}

// ExportExecutionListResult result of listing exports execution history of a export by name
type ExportExecutionListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of export executions.
	Value *[]ExportExecution `json:"value,omitempty"`
}

// ExportExecutionProperties the properties of the export execution.
type ExportExecutionProperties struct {
	// ExecutionType - The type of the export execution. Possible values include: 'OnDemand', 'Scheduled'
	ExecutionType ExecutionType `json:"executionType,omitempty"`
	// Status - The status of the export execution. Possible values include: 'Queued', 'InProgress', 'Completed', 'Failed', 'Timeout', 'NewDataNotAvailable', 'DataNotAvailable'
	Status ExecutionStatus `json:"status,omitempty"`
	// SubmittedBy - The identifier for the entity that executed the export. For OnDemand executions, it is the email id. For Scheduled executions, it is the constant value - System.
	SubmittedBy *string `json:"submittedBy,omitempty"`
	// SubmittedTime - The time when export was queued to be executed.
	SubmittedTime *date.Time `json:"submittedTime,omitempty"`
	// ProcessingStartTime - The time when export was picked up to be executed.
	ProcessingStartTime *date.Time `json:"processingStartTime,omitempty"`
	// ProcessingEndTime - The time when export execution finished.
	ProcessingEndTime *date.Time `json:"processingEndTime,omitempty"`
	// FileName - The name of the file export got written to.
	FileName    *string                 `json:"fileName,omitempty"`
	RunSettings *CommonExportProperties `json:"runSettings,omitempty"`
}

// ExportListResult result of listing exports. It contains a list of available exports in the scope
// provided.
type ExportListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of exports.
	Value *[]Export `json:"value,omitempty"`
}

// ExportProperties the properties of the export.
type ExportProperties struct {
	// Schedule - Has schedule information for the export.
	Schedule *ExportSchedule `json:"schedule,omitempty"`
	// Format - The format of the export being delivered. Possible values include: 'Csv'
	Format FormatType `json:"format,omitempty"`
	// DeliveryInfo - Has delivery information for the export.
	DeliveryInfo *ExportDeliveryInfo `json:"deliveryInfo,omitempty"`
	// Definition - Has definition for the export.
	Definition *QueryDefinition `json:"definition,omitempty"`
}

// ExportRecurrencePeriod the start and end date for recurrence schedule.
type ExportRecurrencePeriod struct {
	// From - The start date of recurrence.
	From *date.Time `json:"from,omitempty"`
	// To - The end date of recurrence.
	To *date.Time `json:"to,omitempty"`
}

// ExportSchedule the schedule associated with a export.
type ExportSchedule struct {
	// Status - The status of the schedule. Whether active or not. If inactive, the export's scheduled execution is paused. Possible values include: 'Active', 'Inactive'
	Status StatusType `json:"status,omitempty"`
	// Recurrence - The schedule recurrence. Possible values include: 'RecurrenceTypeDaily', 'RecurrenceTypeWeekly', 'RecurrenceTypeMonthly', 'RecurrenceTypeAnnually'
	Recurrence RecurrenceType `json:"recurrence,omitempty"`
	// RecurrencePeriod - Has start and end date of the recurrence. The start date must be in future. If present, the end date must be greater than start date.
	RecurrencePeriod *ExportRecurrencePeriod `json:"recurrencePeriod,omitempty"`
}

// Operation a Cost management REST API operation.
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft.CostManagement.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed: Dimensions, Query.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of listing cost management operations. It contains a list of operations and a
// URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of cost management operations supported by the Microsoft.CostManagement resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// QueryAggregation the aggregation expression to be used in the query.
type QueryAggregation struct {
	// Name - The name of the column to aggregate.
	Name *string `json:"name,omitempty"`
	// Function - The name of the aggregation function to use.
	Function *string `json:"function,omitempty"`
}

// QueryColumn ...
type QueryColumn struct {
	// Name - The name of column.
	Name *string `json:"name,omitempty"`
	// Type - The type of column.
	Type *string `json:"type,omitempty"`
}

// QueryComparisonExpression the comparison expression to be used in the query.
type QueryComparisonExpression struct {
	// Name - The name of the column to use in comparison.
	Name *string `json:"name,omitempty"`
	// Operator - The operator to use for comparison.
	Operator *string `json:"operator,omitempty"`
	// Values - Array of values to use for comparison
	Values *[]string `json:"values,omitempty"`
}

// QueryDataset the definition of data present in the query.
type QueryDataset struct {
	// Granularity - The granularity of rows in the query. Possible values include: 'Daily', 'Hourly'
	Granularity GranularityType `json:"granularity,omitempty"`
	// Configuration - Has configuration information for the data in the export. The configuration will be ignored if aggregation and grouping are provided.
	Configuration *QueryDatasetConfiguration `json:"configuration,omitempty"`
	// Aggregation - Dictionary of aggregation expression to use in the query. The key of each item in the dictionary is the alias for the aggregated column. Query can have up to 2 aggregation clauses.
	Aggregation map[string]*QueryAggregation `json:"aggregation"`
	// Grouping - Array of group by expression to use in the query. Query can have up to 2 group by clauses.
	Grouping *[]QueryGrouping `json:"grouping,omitempty"`
	// Sorting - Array of sorting by columns in query.
	Sorting *[]QuerySortingConfiguration `json:"sorting,omitempty"`
	// Filter - Has filter expression to use in the query.
	Filter *QueryFilter `json:"filter,omitempty"`
}

// MarshalJSON is the custom marshaler for QueryDataset.
func (qd QueryDataset) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if qd.Granularity != "" {
		objectMap["granularity"] = qd.Granularity
	}
	if qd.Configuration != nil {
		objectMap["configuration"] = qd.Configuration
	}
	if qd.Aggregation != nil {
		objectMap["aggregation"] = qd.Aggregation
	}
	if qd.Grouping != nil {
		objectMap["grouping"] = qd.Grouping
	}
	if qd.Sorting != nil {
		objectMap["sorting"] = qd.Sorting
	}
	if qd.Filter != nil {
		objectMap["filter"] = qd.Filter
	}
	return json.Marshal(objectMap)
}

// QueryDatasetConfiguration the configuration of dataset in the query.
type QueryDatasetConfiguration struct {
	// Columns - Array of column names to be included in the query. Any valid query column name is allowed. If not provided, then query includes all columns.
	Columns *[]string `json:"columns,omitempty"`
}

// QueryDefinition the definition of a query.
type QueryDefinition struct {
	// Type - The type of the query.
	Type *string `json:"type,omitempty"`
	// Timeframe - The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: 'WeekToDate', 'MonthToDate', 'YearToDate', 'TheLastWeek', 'TheLastMonth', 'TheLastYear', 'Custom'
	Timeframe TimeframeType `json:"timeframe,omitempty"`
	// TimePeriod - Has time period for pulling data for the query.
	TimePeriod *QueryTimePeriod `json:"timePeriod,omitempty"`
	// Dataset - Has definition for data in this query.
	Dataset *QueryDataset `json:"dataset,omitempty"`
}

// QueryFilter the filter expression to be used in the export.
type QueryFilter struct {
	// And - The logical "AND" expression. Must have at least 2 items.
	And *[]QueryFilter `json:"and,omitempty"`
	// Or - The logical "OR" expression. Must have at least 2 items.
	Or *[]QueryFilter `json:"or,omitempty"`
	// Not - The logical "NOT" expression.
	Not *QueryFilter `json:"not,omitempty"`
	// Dimension - Has comparison expression for a dimension
	Dimension *QueryComparisonExpression `json:"dimension,omitempty"`
	// Tag - Has comparison expression for a tag
	Tag *QueryComparisonExpression `json:"tag,omitempty"`
}

// QueryGrouping the group by expression to be used in the query.
type QueryGrouping struct {
	// Type - Has type of the column to group. Possible values include: 'QueryColumnTypeTag', 'QueryColumnTypeDimension'
	Type QueryColumnType `json:"type,omitempty"`
	// Name - The name of the column to group.
	Name *string `json:"name,omitempty"`
}

// QueryProperties ...
type QueryProperties struct {
	// NextLink - The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Columns - Array of columns
	Columns *[]QueryColumn `json:"columns,omitempty"`
	// Rows - Array of rows
	Rows *[][]interface{} `json:"rows,omitempty"`
}

// QueryResult result of query. It contains all columns listed under groupings and aggregation.
type QueryResult struct {
	autorest.Response `json:"-"`
	*QueryProperties  `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for QueryResult.
func (qr QueryResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if qr.QueryProperties != nil {
		objectMap["properties"] = qr.QueryProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for QueryResult struct.
func (qr *QueryResult) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var queryProperties QueryProperties
				err = json.Unmarshal(*v, &queryProperties)
				if err != nil {
					return err
				}
				qr.QueryProperties = &queryProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				qr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				qr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				qr.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				qr.Tags = tags
			}
		}
	}

	return nil
}

// QuerySortingConfiguration the configuration for sorting in the query.
type QuerySortingConfiguration struct {
	// QuerySortingDirection - The sorting direction. Possible values include: 'Ascending', 'Descending'
	QuerySortingDirection SortDirection `json:"querySortingDirection,omitempty"`
	// Name - The name of the column to use in sorting.
	Name *string `json:"name,omitempty"`
}

// QueryTimePeriod the start and end date for pulling data for the query.
type QueryTimePeriod struct {
	// From - The start date to pull data from.
	From *date.Time `json:"from,omitempty"`
	// To - The end date to pull data to.
	To *date.Time `json:"to,omitempty"`
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}
