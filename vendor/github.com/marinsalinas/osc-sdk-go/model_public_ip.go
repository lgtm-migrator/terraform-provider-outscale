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

// PublicIp Information about the public IP.
type PublicIp struct {
	// (Required in a Net) The ID representing the association of the EIP with the VM or the NIC.
	LinkPublicIpId *string `json:"LinkPublicIpId,omitempty"`
	// The account ID of the owner of the NIC.
	NicAccountId *string `json:"NicAccountId,omitempty"`
	// The ID of the NIC the EIP is associated with (if any).
	NicId *string `json:"NicId,omitempty"`
	// The private IP address associated with the EIP.
	PrivateIp *string `json:"PrivateIp,omitempty"`
	// The External IP address (EIP) associated with the NAT service.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The allocation ID of the EIP associated with the NAT service.
	PublicIpId *string `json:"PublicIpId,omitempty"`
	// One or more tags associated with the EIP.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the VM the External IP (EIP) is associated with (if any).
	VmId *string `json:"VmId,omitempty"`
}

// GetLinkPublicIpId returns the LinkPublicIpId field value if set, zero value otherwise.
func (o *PublicIp) GetLinkPublicIpId() string {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret
	}
	return *o.LinkPublicIpId
}

// GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetLinkPublicIpIdOk() (string, bool) {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret, false
	}
	return *o.LinkPublicIpId, true
}

// HasLinkPublicIpId returns a boolean if a field has been set.
func (o *PublicIp) HasLinkPublicIpId() bool {
	if o != nil && o.LinkPublicIpId != nil {
		return true
	}

	return false
}

// SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.
func (o *PublicIp) SetLinkPublicIpId(v string) {
	o.LinkPublicIpId = &v
}

// GetNicAccountId returns the NicAccountId field value if set, zero value otherwise.
func (o *PublicIp) GetNicAccountId() string {
	if o == nil || o.NicAccountId == nil {
		var ret string
		return ret
	}
	return *o.NicAccountId
}

// GetNicAccountIdOk returns a tuple with the NicAccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetNicAccountIdOk() (string, bool) {
	if o == nil || o.NicAccountId == nil {
		var ret string
		return ret, false
	}
	return *o.NicAccountId, true
}

// HasNicAccountId returns a boolean if a field has been set.
func (o *PublicIp) HasNicAccountId() bool {
	if o != nil && o.NicAccountId != nil {
		return true
	}

	return false
}

// SetNicAccountId gets a reference to the given string and assigns it to the NicAccountId field.
func (o *PublicIp) SetNicAccountId(v string) {
	o.NicAccountId = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *PublicIp) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetNicIdOk() (string, bool) {
	if o == nil || o.NicId == nil {
		var ret string
		return ret, false
	}
	return *o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *PublicIp) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *PublicIp) SetNicId(v string) {
	o.NicId = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *PublicIp) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetPrivateIpOk() (string, bool) {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret, false
	}
	return *o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *PublicIp) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *PublicIp) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *PublicIp) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetPublicIpOk() (string, bool) {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *PublicIp) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *PublicIp) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpId returns the PublicIpId field value if set, zero value otherwise.
func (o *PublicIp) GetPublicIpId() string {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetPublicIpIdOk() (string, bool) {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIpId, true
}

// HasPublicIpId returns a boolean if a field has been set.
func (o *PublicIp) HasPublicIpId() bool {
	if o != nil && o.PublicIpId != nil {
		return true
	}

	return false
}

// SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.
func (o *PublicIp) SetPublicIpId(v string) {
	o.PublicIpId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PublicIp) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetTagsOk() ([]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PublicIp) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *PublicIp) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *PublicIp) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetVmIdOk() (string, bool) {
	if o == nil || o.VmId == nil {
		var ret string
		return ret, false
	}
	return *o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *PublicIp) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *PublicIp) SetVmId(v string) {
	o.VmId = &v
}

type NullablePublicIp struct {
	Value        PublicIp
	ExplicitNull bool
}

func (v NullablePublicIp) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullablePublicIp) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
