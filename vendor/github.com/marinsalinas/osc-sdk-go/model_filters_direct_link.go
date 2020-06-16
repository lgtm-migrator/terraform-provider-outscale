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

// FiltersDirectLink One or more filters.
type FiltersDirectLink struct {
	// The IDs of the DirectLinks.
	DirectLinkIds *[]string `json:"DirectLinkIds,omitempty"`
}

// GetDirectLinkIds returns the DirectLinkIds field value if set, zero value otherwise.
func (o *FiltersDirectLink) GetDirectLinkIds() []string {
	if o == nil || o.DirectLinkIds == nil {
		var ret []string
		return ret
	}
	return *o.DirectLinkIds
}

// GetDirectLinkIdsOk returns a tuple with the DirectLinkIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDirectLink) GetDirectLinkIdsOk() ([]string, bool) {
	if o == nil || o.DirectLinkIds == nil {
		var ret []string
		return ret, false
	}
	return *o.DirectLinkIds, true
}

// HasDirectLinkIds returns a boolean if a field has been set.
func (o *FiltersDirectLink) HasDirectLinkIds() bool {
	if o != nil && o.DirectLinkIds != nil {
		return true
	}

	return false
}

// SetDirectLinkIds gets a reference to the given []string and assigns it to the DirectLinkIds field.
func (o *FiltersDirectLink) SetDirectLinkIds(v []string) {
	o.DirectLinkIds = &v
}

type NullableFiltersDirectLink struct {
	Value        FiltersDirectLink
	ExplicitNull bool
}

func (v NullableFiltersDirectLink) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersDirectLink) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
