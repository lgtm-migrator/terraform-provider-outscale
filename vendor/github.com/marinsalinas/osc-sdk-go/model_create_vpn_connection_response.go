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

// CreateVpnConnectionResponse struct for CreateVpnConnectionResponse
type CreateVpnConnectionResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	VpnConnection   *VpnConnection   `json:"VpnConnection,omitempty"`
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateVpnConnectionResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateVpnConnectionResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateVpnConnectionResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVpnConnection returns the VpnConnection field value if set, zero value otherwise.
func (o *CreateVpnConnectionResponse) GetVpnConnection() VpnConnection {
	if o == nil || o.VpnConnection == nil {
		var ret VpnConnection
		return ret
	}
	return *o.VpnConnection
}

// GetVpnConnectionOk returns a tuple with the VpnConnection field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionResponse) GetVpnConnectionOk() (VpnConnection, bool) {
	if o == nil || o.VpnConnection == nil {
		var ret VpnConnection
		return ret, false
	}
	return *o.VpnConnection, true
}

// HasVpnConnection returns a boolean if a field has been set.
func (o *CreateVpnConnectionResponse) HasVpnConnection() bool {
	if o != nil && o.VpnConnection != nil {
		return true
	}

	return false
}

// SetVpnConnection gets a reference to the given VpnConnection and assigns it to the VpnConnection field.
func (o *CreateVpnConnectionResponse) SetVpnConnection(v VpnConnection) {
	o.VpnConnection = &v
}

type NullableCreateVpnConnectionResponse struct {
	Value        CreateVpnConnectionResponse
	ExplicitNull bool
}

func (v NullableCreateVpnConnectionResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateVpnConnectionResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
