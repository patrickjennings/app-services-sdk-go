/*
 * Service Accounts API Documentation
 *
 * This is the API documentation for Service Accounts
 *
 * API version: 1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// InlineResponse400ResponseMediaType struct for InlineResponse400ResponseMediaType
type InlineResponse400ResponseMediaType struct {

	Type *string `json:"type,omitempty"`

	Subtype *string `json:"subtype,omitempty"`

	Parameters *map[string]string `json:"parameters,omitempty"`

	WildcardType *bool `json:"wildcardType,omitempty"`

	WildcardSubtype *bool `json:"wildcardSubtype,omitempty"`

}

// NewInlineResponse400ResponseMediaType instantiates a new InlineResponse400ResponseMediaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400ResponseMediaType() *InlineResponse400ResponseMediaType {
	this := InlineResponse400ResponseMediaType{}
	return &this
}

// NewInlineResponse400ResponseMediaTypeWithDefaults instantiates a new InlineResponse400ResponseMediaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400ResponseMediaTypeWithDefaults() *InlineResponse400ResponseMediaType {
	this := InlineResponse400ResponseMediaType{}






	return &this
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse400ResponseMediaType) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseMediaType) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse400ResponseMediaType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse400ResponseMediaType) SetType(v string) {
	o.Type = &v
}


// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *InlineResponse400ResponseMediaType) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseMediaType) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *InlineResponse400ResponseMediaType) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *InlineResponse400ResponseMediaType) SetSubtype(v string) {
	o.Subtype = &v
}


// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *InlineResponse400ResponseMediaType) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseMediaType) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *InlineResponse400ResponseMediaType) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *InlineResponse400ResponseMediaType) SetParameters(v map[string]string) {
	o.Parameters = &v
}


// GetWildcardType returns the WildcardType field value if set, zero value otherwise.
func (o *InlineResponse400ResponseMediaType) GetWildcardType() bool {
	if o == nil || o.WildcardType == nil {
		var ret bool
		return ret
	}
	return *o.WildcardType
}

// GetWildcardTypeOk returns a tuple with the WildcardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseMediaType) GetWildcardTypeOk() (*bool, bool) {
	if o == nil || o.WildcardType == nil {
		return nil, false
	}
	return o.WildcardType, true
}

// HasWildcardType returns a boolean if a field has been set.
func (o *InlineResponse400ResponseMediaType) HasWildcardType() bool {
	if o != nil && o.WildcardType != nil {
		return true
	}

	return false
}

// SetWildcardType gets a reference to the given bool and assigns it to the WildcardType field.
func (o *InlineResponse400ResponseMediaType) SetWildcardType(v bool) {
	o.WildcardType = &v
}


// GetWildcardSubtype returns the WildcardSubtype field value if set, zero value otherwise.
func (o *InlineResponse400ResponseMediaType) GetWildcardSubtype() bool {
	if o == nil || o.WildcardSubtype == nil {
		var ret bool
		return ret
	}
	return *o.WildcardSubtype
}

// GetWildcardSubtypeOk returns a tuple with the WildcardSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseMediaType) GetWildcardSubtypeOk() (*bool, bool) {
	if o == nil || o.WildcardSubtype == nil {
		return nil, false
	}
	return o.WildcardSubtype, true
}

// HasWildcardSubtype returns a boolean if a field has been set.
func (o *InlineResponse400ResponseMediaType) HasWildcardSubtype() bool {
	if o != nil && o.WildcardSubtype != nil {
		return true
	}

	return false
}

// SetWildcardSubtype gets a reference to the given bool and assigns it to the WildcardSubtype field.
func (o *InlineResponse400ResponseMediaType) SetWildcardSubtype(v bool) {
	o.WildcardSubtype = &v
}


func (o InlineResponse400ResponseMediaType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
    
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
    
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
    
	if o.WildcardType != nil {
		toSerialize["wildcardType"] = o.WildcardType
	}
    
	if o.WildcardSubtype != nil {
		toSerialize["wildcardSubtype"] = o.WildcardSubtype
	}
    
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400ResponseMediaType struct {
	value *InlineResponse400ResponseMediaType
	isSet bool
}

func (v NullableInlineResponse400ResponseMediaType) Get() *InlineResponse400ResponseMediaType {
	return v.value
}

func (v *NullableInlineResponse400ResponseMediaType) Set(val *InlineResponse400ResponseMediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400ResponseMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400ResponseMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400ResponseMediaType(val *InlineResponse400ResponseMediaType) *NullableInlineResponse400ResponseMediaType {
	return &NullableInlineResponse400ResponseMediaType{value: val, isSet: true}
}

func (v NullableInlineResponse400ResponseMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400ResponseMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

