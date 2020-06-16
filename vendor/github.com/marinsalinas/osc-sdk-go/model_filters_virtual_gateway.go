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

// FiltersVirtualGateway One or more filters.
type FiltersVirtualGateway struct {
	// The types of the virtual gateways (only `ipsec.1` is supported).
	ConnectionTypes *[]string `json:"ConnectionTypes,omitempty"`
	// The IDs of the Nets the virtual gateways are attached to.
	LinkNetIds *[]string `json:"LinkNetIds,omitempty"`
	// The current states of the attachments between the virtual gateways and the Nets (`attaching` \\| `attached` \\| `detaching` \\| `detached`).
	LinkStates *[]string `json:"LinkStates,omitempty"`
	// The states of the virtual gateways (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States *[]string `json:"States,omitempty"`
	// The keys of the tags associated with the virtual gateways.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the virtual gateways.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the virtual gateways, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The IDs of the virtual gateways.
	VirtualGatewayIds *[]string `json:"VirtualGatewayIds,omitempty"`
}

// GetConnectionTypes returns the ConnectionTypes field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetConnectionTypes() []string {
	if o == nil || o.ConnectionTypes == nil {
		var ret []string
		return ret
	}
	return *o.ConnectionTypes
}

// GetConnectionTypesOk returns a tuple with the ConnectionTypes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetConnectionTypesOk() ([]string, bool) {
	if o == nil || o.ConnectionTypes == nil {
		var ret []string
		return ret, false
	}
	return *o.ConnectionTypes, true
}

// HasConnectionTypes returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasConnectionTypes() bool {
	if o != nil && o.ConnectionTypes != nil {
		return true
	}

	return false
}

// SetConnectionTypes gets a reference to the given []string and assigns it to the ConnectionTypes field.
func (o *FiltersVirtualGateway) SetConnectionTypes(v []string) {
	o.ConnectionTypes = &v
}

// GetLinkNetIds returns the LinkNetIds field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetLinkNetIds() []string {
	if o == nil || o.LinkNetIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkNetIds
}

// GetLinkNetIdsOk returns a tuple with the LinkNetIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetLinkNetIdsOk() ([]string, bool) {
	if o == nil || o.LinkNetIds == nil {
		var ret []string
		return ret, false
	}
	return *o.LinkNetIds, true
}

// HasLinkNetIds returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasLinkNetIds() bool {
	if o != nil && o.LinkNetIds != nil {
		return true
	}

	return false
}

// SetLinkNetIds gets a reference to the given []string and assigns it to the LinkNetIds field.
func (o *FiltersVirtualGateway) SetLinkNetIds(v []string) {
	o.LinkNetIds = &v
}

// GetLinkStates returns the LinkStates field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetLinkStates() []string {
	if o == nil || o.LinkStates == nil {
		var ret []string
		return ret
	}
	return *o.LinkStates
}

// GetLinkStatesOk returns a tuple with the LinkStates field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetLinkStatesOk() ([]string, bool) {
	if o == nil || o.LinkStates == nil {
		var ret []string
		return ret, false
	}
	return *o.LinkStates, true
}

// HasLinkStates returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasLinkStates() bool {
	if o != nil && o.LinkStates != nil {
		return true
	}

	return false
}

// SetLinkStates gets a reference to the given []string and assigns it to the LinkStates field.
func (o *FiltersVirtualGateway) SetLinkStates(v []string) {
	o.LinkStates = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetStatesOk() ([]string, bool) {
	if o == nil || o.States == nil {
		var ret []string
		return ret, false
	}
	return *o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersVirtualGateway) SetStates(v []string) {
	o.States = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetTagKeysOk() ([]string, bool) {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret, false
	}
	return *o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersVirtualGateway) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetTagValuesOk() ([]string, bool) {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret, false
	}
	return *o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersVirtualGateway) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersVirtualGateway) SetTags(v []string) {
	o.Tags = &v
}

// GetVirtualGatewayIds returns the VirtualGatewayIds field value if set, zero value otherwise.
func (o *FiltersVirtualGateway) GetVirtualGatewayIds() []string {
	if o == nil || o.VirtualGatewayIds == nil {
		var ret []string
		return ret
	}
	return *o.VirtualGatewayIds
}

// GetVirtualGatewayIdsOk returns a tuple with the VirtualGatewayIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVirtualGateway) GetVirtualGatewayIdsOk() ([]string, bool) {
	if o == nil || o.VirtualGatewayIds == nil {
		var ret []string
		return ret, false
	}
	return *o.VirtualGatewayIds, true
}

// HasVirtualGatewayIds returns a boolean if a field has been set.
func (o *FiltersVirtualGateway) HasVirtualGatewayIds() bool {
	if o != nil && o.VirtualGatewayIds != nil {
		return true
	}

	return false
}

// SetVirtualGatewayIds gets a reference to the given []string and assigns it to the VirtualGatewayIds field.
func (o *FiltersVirtualGateway) SetVirtualGatewayIds(v []string) {
	o.VirtualGatewayIds = &v
}

type NullableFiltersVirtualGateway struct {
	Value        FiltersVirtualGateway
	ExplicitNull bool
}

func (v NullableFiltersVirtualGateway) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersVirtualGateway) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
