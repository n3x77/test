/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: 0.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Systemstatus struct for Systemstatus
type Systemstatus struct {
	SystemstatusId *int32 `json:"systemstatus_id,omitempty"`
	SystemstatusName string `json:"systemstatus_name"`
}

// NewSystemstatus instantiates a new Systemstatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemstatus(systemstatusName string, ) *Systemstatus {
	this := Systemstatus{}
	this.SystemstatusName = systemstatusName
	return &this
}

// NewSystemstatusWithDefaults instantiates a new Systemstatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemstatusWithDefaults() *Systemstatus {
	this := Systemstatus{}
	return &this
}

// GetSystemstatusId returns the SystemstatusId field value if set, zero value otherwise.
func (o *Systemstatus) GetSystemstatusId() int32 {
	if o == nil || o.SystemstatusId == nil {
		var ret int32
		return ret
	}
	return *o.SystemstatusId
}

// GetSystemstatusIdOk returns a tuple with the SystemstatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemstatus) GetSystemstatusIdOk() (*int32, bool) {
	if o == nil || o.SystemstatusId == nil {
		return nil, false
	}
	return o.SystemstatusId, true
}

// HasSystemstatusId returns a boolean if a field has been set.
func (o *Systemstatus) HasSystemstatusId() bool {
	if o != nil && o.SystemstatusId != nil {
		return true
	}

	return false
}

// SetSystemstatusId gets a reference to the given int32 and assigns it to the SystemstatusId field.
func (o *Systemstatus) SetSystemstatusId(v int32) {
	o.SystemstatusId = &v
}

// GetSystemstatusName returns the SystemstatusName field value
func (o *Systemstatus) GetSystemstatusName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SystemstatusName
}

// GetSystemstatusNameOk returns a tuple with the SystemstatusName field value
// and a boolean to check if the value has been set.
func (o *Systemstatus) GetSystemstatusNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemstatusName, true
}

// SetSystemstatusName sets field value
func (o *Systemstatus) SetSystemstatusName(v string) {
	o.SystemstatusName = v
}

func (o Systemstatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemstatusId != nil {
		toSerialize["systemstatus_id"] = o.SystemstatusId
	}
	if true {
		toSerialize["systemstatus_name"] = o.SystemstatusName
	}
	return json.Marshal(toSerialize)
}

type NullableSystemstatus struct {
	value *Systemstatus
	isSet bool
}

func (v NullableSystemstatus) Get() *Systemstatus {
	return v.value
}

func (v *NullableSystemstatus) Set(val *Systemstatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemstatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemstatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemstatus(val *Systemstatus) *NullableSystemstatus {
	return &NullableSystemstatus{value: val, isSet: true}
}

func (v NullableSystemstatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemstatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

