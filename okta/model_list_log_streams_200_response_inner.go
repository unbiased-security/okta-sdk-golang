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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

//model_oneof.mustache
// ListLogStreams200ResponseInner - struct for ListLogStreams200ResponseInner
type ListLogStreams200ResponseInner struct {
	LogStreamAws *LogStreamAws
	LogStreamSplunk *LogStreamSplunk
}

// LogStreamAwsAsListLogStreams200ResponseInner is a convenience function that returns LogStreamAws wrapped in ListLogStreams200ResponseInner
func LogStreamAwsAsListLogStreams200ResponseInner(v *LogStreamAws) ListLogStreams200ResponseInner {
	return ListLogStreams200ResponseInner{
		LogStreamAws: v,
	}
}

// LogStreamSplunkAsListLogStreams200ResponseInner is a convenience function that returns LogStreamSplunk wrapped in ListLogStreams200ResponseInner
func LogStreamSplunkAsListLogStreams200ResponseInner(v *LogStreamSplunk) ListLogStreams200ResponseInner {
	return ListLogStreams200ResponseInner{
		LogStreamSplunk: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListLogStreams200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'LogStreamAws'
	if jsonDict["type"] == "LogStreamAws" {
		// try to unmarshal JSON data into LogStreamAws
		err = json.Unmarshal(data, &dst.LogStreamAws)
		if err == nil {
			return nil // data stored in dst.LogStreamAws, return on the first match
		} else {
			dst.LogStreamAws = nil
			return fmt.Errorf("Failed to unmarshal ListLogStreams200ResponseInner as LogStreamAws: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LogStreamSplunk'
	if jsonDict["type"] == "LogStreamSplunk" {
		// try to unmarshal JSON data into LogStreamSplunk
		err = json.Unmarshal(data, &dst.LogStreamSplunk)
		if err == nil {
			return nil // data stored in dst.LogStreamSplunk, return on the first match
		} else {
			dst.LogStreamSplunk = nil
			return fmt.Errorf("Failed to unmarshal ListLogStreams200ResponseInner as LogStreamSplunk: %s", err.Error())
		}
	}

	// check if the discriminator value is 'aws_eventbridge'
	if jsonDict["type"] == "aws_eventbridge" {
		// try to unmarshal JSON data into LogStreamAws
		err = json.Unmarshal(data, &dst.LogStreamAws)
		if err == nil {
			return nil // data stored in dst.LogStreamAws, return on the first match
		} else {
			dst.LogStreamAws = nil
			return fmt.Errorf("Failed to unmarshal ListLogStreams200ResponseInner as LogStreamAws: %s", err.Error())
		}
	}

	// check if the discriminator value is 'splunk_cloud_logstreaming'
	if jsonDict["type"] == "splunk_cloud_logstreaming" {
		// try to unmarshal JSON data into LogStreamSplunk
		err = json.Unmarshal(data, &dst.LogStreamSplunk)
		if err == nil {
			return nil // data stored in dst.LogStreamSplunk, return on the first match
		} else {
			dst.LogStreamSplunk = nil
			return fmt.Errorf("Failed to unmarshal ListLogStreams200ResponseInner as LogStreamSplunk: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListLogStreams200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.LogStreamAws != nil {
		return json.Marshal(&src.LogStreamAws)
	}

	if src.LogStreamSplunk != nil {
		return json.Marshal(&src.LogStreamSplunk)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListLogStreams200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LogStreamAws != nil {
		return obj.LogStreamAws
	}

	if obj.LogStreamSplunk != nil {
		return obj.LogStreamSplunk
	}

	// all schemas are nil
	return nil
}

type NullableListLogStreams200ResponseInner struct {
	value *ListLogStreams200ResponseInner
	isSet bool
}

func (v NullableListLogStreams200ResponseInner) Get() *ListLogStreams200ResponseInner {
	return v.value
}

func (v *NullableListLogStreams200ResponseInner) Set(val *ListLogStreams200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListLogStreams200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListLogStreams200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLogStreams200ResponseInner(val *ListLogStreams200ResponseInner) *NullableListLogStreams200ResponseInner {
	return &NullableListLogStreams200ResponseInner{value: val, isSet: true}
}

func (v NullableListLogStreams200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLogStreams200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

