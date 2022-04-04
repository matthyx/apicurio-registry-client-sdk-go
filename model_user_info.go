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

// UserInfo Information about a single user.
type UserInfo struct {
	Username *string `json:"username,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Admin *bool `json:"admin,omitempty"`
	Developer *bool `json:"developer,omitempty"`
	Viewer *bool `json:"viewer,omitempty"`
}

// NewUserInfo instantiates a new UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfo() *UserInfo {
	this := UserInfo{}
	return &this
}

// NewUserInfoWithDefaults instantiates a new UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoWithDefaults() *UserInfo {
	this := UserInfo{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserInfo) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserInfo) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserInfo) SetUsername(v string) {
	o.Username = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *UserInfo) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *UserInfo) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *UserInfo) SetAdmin(v bool) {
	o.Admin = &v
}

// GetDeveloper returns the Developer field value if set, zero value otherwise.
func (o *UserInfo) GetDeveloper() bool {
	if o == nil || o.Developer == nil {
		var ret bool
		return ret
	}
	return *o.Developer
}

// GetDeveloperOk returns a tuple with the Developer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetDeveloperOk() (*bool, bool) {
	if o == nil || o.Developer == nil {
		return nil, false
	}
	return o.Developer, true
}

// HasDeveloper returns a boolean if a field has been set.
func (o *UserInfo) HasDeveloper() bool {
	if o != nil && o.Developer != nil {
		return true
	}

	return false
}

// SetDeveloper gets a reference to the given bool and assigns it to the Developer field.
func (o *UserInfo) SetDeveloper(v bool) {
	o.Developer = &v
}

// GetViewer returns the Viewer field value if set, zero value otherwise.
func (o *UserInfo) GetViewer() bool {
	if o == nil || o.Viewer == nil {
		var ret bool
		return ret
	}
	return *o.Viewer
}

// GetViewerOk returns a tuple with the Viewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetViewerOk() (*bool, bool) {
	if o == nil || o.Viewer == nil {
		return nil, false
	}
	return o.Viewer, true
}

// HasViewer returns a boolean if a field has been set.
func (o *UserInfo) HasViewer() bool {
	if o != nil && o.Viewer != nil {
		return true
	}

	return false
}

// SetViewer gets a reference to the given bool and assigns it to the Viewer field.
func (o *UserInfo) SetViewer(v bool) {
	o.Viewer = &v
}

func (o UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.Developer != nil {
		toSerialize["developer"] = o.Developer
	}
	if o.Viewer != nil {
		toSerialize["viewer"] = o.Viewer
	}
	return json.Marshal(toSerialize)
}

type NullableUserInfo struct {
	value *UserInfo
	isSet bool
}

func (v NullableUserInfo) Get() *UserInfo {
	return v.value
}

func (v *NullableUserInfo) Set(val *UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfo(val *UserInfo) *NullableUserInfo {
	return &NullableUserInfo{value: val, isSet: true}
}

func (v NullableUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


