/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// PasswordPolicyRule struct for PasswordPolicyRule
type PasswordPolicyRule struct {
	PolicyRule
	Actions              *PasswordPolicyRuleActions    `json:"actions,omitempty"`
	Conditions           *PasswordPolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRule PasswordPolicyRule

// NewPasswordPolicyRule instantiates a new PasswordPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRule() *PasswordPolicyRule {
	this := PasswordPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewPasswordPolicyRuleWithDefaults instantiates a new PasswordPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRuleWithDefaults() *PasswordPolicyRule {
	this := PasswordPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *PasswordPolicyRule) GetActions() PasswordPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret PasswordPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRule) GetActionsOk() (*PasswordPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *PasswordPolicyRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given PasswordPolicyRuleActions and assigns it to the Actions field.
func (o *PasswordPolicyRule) SetActions(v PasswordPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PasswordPolicyRule) GetConditions() PasswordPolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PasswordPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRule) GetConditionsOk() (*PasswordPolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PasswordPolicyRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PasswordPolicyRuleConditions and assigns it to the Conditions field.
func (o *PasswordPolicyRule) SetConditions(v PasswordPolicyRuleConditions) {
	o.Conditions = &v
}

func (o PasswordPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyRule, errPolicyRule := json.Marshal(o.PolicyRule)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	errPolicyRule = json.Unmarshal([]byte(serializedPolicyRule), &toSerialize)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	type PasswordPolicyRuleWithoutEmbeddedStruct struct {
		Actions    *PasswordPolicyRuleActions    `json:"actions,omitempty"`
		Conditions *PasswordPolicyRuleConditions `json:"conditions,omitempty"`
	}

	varPasswordPolicyRuleWithoutEmbeddedStruct := PasswordPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPasswordPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varPasswordPolicyRule := _PasswordPolicyRule{}
		varPasswordPolicyRule.Actions = varPasswordPolicyRuleWithoutEmbeddedStruct.Actions
		varPasswordPolicyRule.Conditions = varPasswordPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = PasswordPolicyRule(varPasswordPolicyRule)
	} else {
		return err
	}

	varPasswordPolicyRule := _PasswordPolicyRule{}

	err = json.Unmarshal(bytes, &varPasswordPolicyRule)
	if err == nil {
		o.PolicyRule = varPasswordPolicyRule.PolicyRule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")

		// remove fields from embedded structs
		reflectPolicyRule := reflect.ValueOf(o.PolicyRule)
		for i := 0; i < reflectPolicyRule.Type().NumField(); i++ {
			t := reflectPolicyRule.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyRule struct {
	value *PasswordPolicyRule
	isSet bool
}

func (v NullablePasswordPolicyRule) Get() *PasswordPolicyRule {
	return v.value
}

func (v *NullablePasswordPolicyRule) Set(val *PasswordPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRule(val *PasswordPolicyRule) *NullablePasswordPolicyRule {
	return &NullablePasswordPolicyRule{value: val, isSet: true}
}

func (v NullablePasswordPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
