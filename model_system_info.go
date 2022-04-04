/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.3-SNAPSHOT
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
)

// SystemInfo struct for SystemInfo
type SystemInfo struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
	BuiltOn *string `json:"builtOn,omitempty"`
}

// NewSystemInfo instantiates a new SystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfo() *SystemInfo {
	this := SystemInfo{}
	return &this
}

// NewSystemInfoWithDefaults instantiates a new SystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoWithDefaults() *SystemInfo {
	this := SystemInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemInfo) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SystemInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SystemInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SystemInfo) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SystemInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SystemInfo) SetVersion(v string) {
	o.Version = &v
}

// GetBuiltOn returns the BuiltOn field value if set, zero value otherwise.
func (o *SystemInfo) GetBuiltOn() string {
	if o == nil || o.BuiltOn == nil {
		var ret string
		return ret
	}
	return *o.BuiltOn
}

// GetBuiltOnOk returns a tuple with the BuiltOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetBuiltOnOk() (*string, bool) {
	if o == nil || o.BuiltOn == nil {
		return nil, false
	}
	return o.BuiltOn, true
}

// HasBuiltOn returns a boolean if a field has been set.
func (o *SystemInfo) HasBuiltOn() bool {
	if o != nil && o.BuiltOn != nil {
		return true
	}

	return false
}

// SetBuiltOn gets a reference to the given string and assigns it to the BuiltOn field.
func (o *SystemInfo) SetBuiltOn(v string) {
	o.BuiltOn = &v
}

func (o SystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.BuiltOn != nil {
		toSerialize["builtOn"] = o.BuiltOn
	}
	return json.Marshal(toSerialize)
}

type NullableSystemInfo struct {
	value *SystemInfo
	isSet bool
}

func (v NullableSystemInfo) Get() *SystemInfo {
	return v.value
}

func (v *NullableSystemInfo) Set(val *SystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfo(val *SystemInfo) *NullableSystemInfo {
	return &NullableSystemInfo{value: val, isSet: true}
}

func (v NullableSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


