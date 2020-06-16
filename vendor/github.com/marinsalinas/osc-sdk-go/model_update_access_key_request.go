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

// UpdateAccessKeyRequest struct for UpdateAccessKeyRequest
type UpdateAccessKeyRequest struct {
	// The ID of the access key.
	AccessKeyId string `json:"AccessKeyId"`
	// The new state for the access key (`active` \\| `inactive`).
	State string `json:"State"`
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *UpdateAccessKeyRequest) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// SetAccessKeyId sets field value
func (o *UpdateAccessKeyRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetState returns the State field value
func (o *UpdateAccessKeyRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *UpdateAccessKeyRequest) SetState(v string) {
	o.State = v
}

type NullableUpdateAccessKeyRequest struct {
	Value        UpdateAccessKeyRequest
	ExplicitNull bool
}

func (v NullableUpdateAccessKeyRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateAccessKeyRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
