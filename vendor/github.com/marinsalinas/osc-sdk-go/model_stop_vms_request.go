/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// StopVmsRequest struct for StopVmsRequest
type StopVmsRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// Forces the VM to stop.
	ForceStop *bool `json:"ForceStop,omitempty"`
	// One or more IDs of VMs.
	VmIds []string `json:"VmIds"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *StopVmsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StopVmsRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
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

// GetForceStopOk returns a tuple with the ForceStop field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StopVmsRequest) GetForceStopOk() (bool, bool) {
	if o == nil || o.ForceStop == nil {
		var ret bool
		return ret, false
	}
	return *o.ForceStop, true
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

// SetVmIds sets field value
func (o *StopVmsRequest) SetVmIds(v []string) {
	o.VmIds = v
}

type NullableStopVmsRequest struct {
	Value        StopVmsRequest
	ExplicitNull bool
}

func (v NullableStopVmsRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableStopVmsRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
