/*
Leaseweb API for Abuse Handling

This API provides ways to manage the abuse reports you might receive from Leaseweb. To use this API, please request access via your account manager and/or compliance officer. **LIMITED ACCESS** 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package abuse

import (
	"encoding/json"
)

// checks if the GetReportListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportListResult{}

// GetReportListResult struct for GetReportListResult
type GetReportListResult struct {
	// An array of abuse reports.
	Reports []Report `json:"reports,omitempty"`
	Metadata *Metadata `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetReportListResult GetReportListResult

// NewGetReportListResult instantiates a new GetReportListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportListResult() *GetReportListResult {
	this := GetReportListResult{}
	return &this
}

// NewGetReportListResultWithDefaults instantiates a new GetReportListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportListResultWithDefaults() *GetReportListResult {
	this := GetReportListResult{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *GetReportListResult) GetReports() []Report {
	if o == nil || IsNil(o.Reports) {
		var ret []Report
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportListResult) GetReportsOk() ([]Report, bool) {
	if o == nil || IsNil(o.Reports) {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *GetReportListResult) HasReports() bool {
	if o != nil && !IsNil(o.Reports) {
		return true
	}

	return false
}

// SetReports gets a reference to the given []Report and assigns it to the Reports field.
func (o *GetReportListResult) SetReports(v []Report) {
	o.Reports = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetReportListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetReportListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetReportListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

func (o GetReportListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReportListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reports) {
		toSerialize["reports"] = o.Reports
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetReportListResult) UnmarshalJSON(data []byte) (err error) {
	varGetReportListResult := _GetReportListResult{}

	err = json.Unmarshal(data, &varGetReportListResult)

	if err != nil {
		return err
	}

	*o = GetReportListResult(varGetReportListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reports")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetReportListResult struct {
	value *GetReportListResult
	isSet bool
}

func (v NullableGetReportListResult) Get() *GetReportListResult {
	return v.value
}

func (v *NullableGetReportListResult) Set(val *GetReportListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportListResult(val *GetReportListResult) *NullableGetReportListResult {
	return &NullableGetReportListResult{value: val, isSet: true}
}

func (v NullableGetReportListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReportListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

