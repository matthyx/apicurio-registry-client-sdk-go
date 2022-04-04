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

// RuleViolationCause struct for RuleViolationCause
type RuleViolationCause struct {
	Description *string `json:"description,omitempty"`
	Context *string `json:"context,omitempty"`
}

// NewRuleViolationCause instantiates a new RuleViolationCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleViolationCause() *RuleViolationCause {
	this := RuleViolationCause{}
	return &this
}

// NewRuleViolationCauseWithDefaults instantiates a new RuleViolationCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleViolationCauseWithDefaults() *RuleViolationCause {
	this := RuleViolationCause{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleViolationCause) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleViolationCause) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleViolationCause) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleViolationCause) SetDescription(v string) {
	o.Description = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RuleViolationCause) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleViolationCause) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RuleViolationCause) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *RuleViolationCause) SetContext(v string) {
	o.Context = &v
}

func (o RuleViolationCause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableRuleViolationCause struct {
	value *RuleViolationCause
	isSet bool
}

func (v NullableRuleViolationCause) Get() *RuleViolationCause {
	return v.value
}

func (v *NullableRuleViolationCause) Set(val *RuleViolationCause) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleViolationCause) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleViolationCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleViolationCause(val *RuleViolationCause) *NullableRuleViolationCause {
	return &NullableRuleViolationCause{value: val, isSet: true}
}

func (v NullableRuleViolationCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleViolationCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


