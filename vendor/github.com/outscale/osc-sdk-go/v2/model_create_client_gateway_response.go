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

// CreateClientGatewayResponse struct for CreateClientGatewayResponse
type CreateClientGatewayResponse struct {
	ClientGateway   *ClientGateway   `json:"ClientGateway,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateClientGatewayResponse instantiates a new CreateClientGatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientGatewayResponse() *CreateClientGatewayResponse {
	this := CreateClientGatewayResponse{}
	return &this
}

// NewCreateClientGatewayResponseWithDefaults instantiates a new CreateClientGatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientGatewayResponseWithDefaults() *CreateClientGatewayResponse {
	this := CreateClientGatewayResponse{}
	return &this
}

// GetClientGateway returns the ClientGateway field value if set, zero value otherwise.
func (o *CreateClientGatewayResponse) GetClientGateway() ClientGateway {
	if o == nil || o.ClientGateway == nil {
		var ret ClientGateway
		return ret
	}
	return *o.ClientGateway
}

// GetClientGatewayOk returns a tuple with the ClientGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientGatewayResponse) GetClientGatewayOk() (*ClientGateway, bool) {
	if o == nil || o.ClientGateway == nil {
		return nil, false
	}
	return o.ClientGateway, true
}

// HasClientGateway returns a boolean if a field has been set.
func (o *CreateClientGatewayResponse) HasClientGateway() bool {
	if o != nil && o.ClientGateway != nil {
		return true
	}

	return false
}

// SetClientGateway gets a reference to the given ClientGateway and assigns it to the ClientGateway field.
func (o *CreateClientGatewayResponse) SetClientGateway(v ClientGateway) {
	o.ClientGateway = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateClientGatewayResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientGatewayResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateClientGatewayResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateClientGatewayResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateClientGatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientGateway != nil {
		toSerialize["ClientGateway"] = o.ClientGateway
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClientGatewayResponse struct {
	value *CreateClientGatewayResponse
	isSet bool
}

func (v NullableCreateClientGatewayResponse) Get() *CreateClientGatewayResponse {
	return v.value
}

func (v *NullableCreateClientGatewayResponse) Set(val *CreateClientGatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClientGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClientGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClientGatewayResponse(val *CreateClientGatewayResponse) *NullableCreateClientGatewayResponse {
	return &NullableCreateClientGatewayResponse{value: val, isSet: true}
}

func (v NullableCreateClientGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClientGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
