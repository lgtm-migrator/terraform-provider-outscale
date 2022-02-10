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

// StopVmsRequest struct for StopVmsRequest
type StopVmsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// Forces the VM to stop.
	ForceStop *bool `json:"ForceStop,omitempty"`
	// One or more IDs of VMs.
	VmIds []string `json:"VmIds"`
}

// NewStopVmsRequest instantiates a new StopVmsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopVmsRequest(vmIds []string) *StopVmsRequest {
	this := StopVmsRequest{}
	this.VmIds = vmIds
	return &this
}

// NewStopVmsRequestWithDefaults instantiates a new StopVmsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopVmsRequestWithDefaults() *StopVmsRequest {
	this := StopVmsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *StopVmsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopVmsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *StopVmsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *StopVmsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetForceStop returns the ForceStop field value if set, zero value otherwise.
func (o *StopVmsRequest) GetForceStop() bool {
	if o == nil || o.ForceStop == nil {
		var ret bool
		return ret
	}
	return *o.ForceStop
}

// GetForceStopOk returns a tuple with the ForceStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopVmsRequest) GetForceStopOk() (*bool, bool) {
	if o == nil || o.ForceStop == nil {
		return nil, false
	}
	return o.ForceStop, true
}

// HasForceStop returns a boolean if a field has been set.
func (o *StopVmsRequest) HasForceStop() bool {
	if o != nil && o.ForceStop != nil {
		return true
	}

	return false
}

// SetForceStop gets a reference to the given bool and assigns it to the ForceStop field.
func (o *StopVmsRequest) SetForceStop(v bool) {
	o.ForceStop = &v
}

// GetVmIds returns the VmIds field value
func (o *StopVmsRequest) GetVmIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value
// and a boolean to check if the value has been set.
func (o *StopVmsRequest) GetVmIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmIds, true
}

// SetVmIds sets field value
func (o *StopVmsRequest) SetVmIds(v []string) {
	o.VmIds = v
}

func (o StopVmsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.ForceStop != nil {
		toSerialize["ForceStop"] = o.ForceStop
	}
	if true {
		toSerialize["VmIds"] = o.VmIds
	}
	return json.Marshal(toSerialize)
}

type NullableStopVmsRequest struct {
	value *StopVmsRequest
	isSet bool
}

func (v NullableStopVmsRequest) Get() *StopVmsRequest {
	return v.value
}

func (v *NullableStopVmsRequest) Set(val *StopVmsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStopVmsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStopVmsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopVmsRequest(val *StopVmsRequest) *NullableStopVmsRequest {
	return &NullableStopVmsRequest{value: val, isSet: true}
}

func (v NullableStopVmsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopVmsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
