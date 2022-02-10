/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.16
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UnlinkRouteTableRequest struct for UnlinkRouteTableRequest
type UnlinkRouteTableRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the association between the route table and the Subnet.
	LinkRouteTableId string `json:"LinkRouteTableId"`
}

// NewUnlinkRouteTableRequest instantiates a new UnlinkRouteTableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkRouteTableRequest(linkRouteTableId string) *UnlinkRouteTableRequest {
	this := UnlinkRouteTableRequest{}
	this.LinkRouteTableId = linkRouteTableId
	return &this
}

// NewUnlinkRouteTableRequestWithDefaults instantiates a new UnlinkRouteTableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkRouteTableRequestWithDefaults() *UnlinkRouteTableRequest {
	this := UnlinkRouteTableRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkRouteTableRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkRouteTableRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkRouteTableRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkRouteTableRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLinkRouteTableId returns the LinkRouteTableId field value
func (o *UnlinkRouteTableRequest) GetLinkRouteTableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkRouteTableId
}

// GetLinkRouteTableIdOk returns a tuple with the LinkRouteTableId field value
// and a boolean to check if the value has been set.
func (o *UnlinkRouteTableRequest) GetLinkRouteTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkRouteTableId, true
}

// SetLinkRouteTableId sets field value
func (o *UnlinkRouteTableRequest) SetLinkRouteTableId(v string) {
	o.LinkRouteTableId = v
}

func (o UnlinkRouteTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["LinkRouteTableId"] = o.LinkRouteTableId
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkRouteTableRequest struct {
	value *UnlinkRouteTableRequest
	isSet bool
}

func (v NullableUnlinkRouteTableRequest) Get() *UnlinkRouteTableRequest {
	return v.value
}

func (v *NullableUnlinkRouteTableRequest) Set(val *UnlinkRouteTableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkRouteTableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkRouteTableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkRouteTableRequest(val *UnlinkRouteTableRequest) *NullableUnlinkRouteTableRequest {
	return &NullableUnlinkRouteTableRequest{value: val, isSet: true}
}

func (v NullableUnlinkRouteTableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkRouteTableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
