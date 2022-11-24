/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.3.2-SNAPSHOT
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
)

// CreateGroupMetaData 
type CreateGroupMetaData struct {
	Description *string `json:"description,omitempty"`
	// User-defined name-value pairs. Name and value must be strings.
	Properties *map[string]string `json:"properties,omitempty"`
	// 
	Id string `json:"id"`
}

// NewCreateGroupMetaData instantiates a new CreateGroupMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupMetaData(id string) *CreateGroupMetaData {
	this := CreateGroupMetaData{}
	this.Id = id
	return &this
}

// NewCreateGroupMetaDataWithDefaults instantiates a new CreateGroupMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupMetaDataWithDefaults() *CreateGroupMetaData {
	this := CreateGroupMetaData{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGroupMetaData) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupMetaData) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGroupMetaData) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGroupMetaData) SetDescription(v string) {
	o.Description = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateGroupMetaData) GetProperties() map[string]string {
	if o == nil || isNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupMetaData) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Properties) {
    return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateGroupMetaData) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *CreateGroupMetaData) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetId returns the Id field value
func (o *CreateGroupMetaData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateGroupMetaData) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateGroupMetaData) SetId(v string) {
	o.Id = v
}

func (o CreateGroupMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGroupMetaData struct {
	value *CreateGroupMetaData
	isSet bool
}

func (v NullableCreateGroupMetaData) Get() *CreateGroupMetaData {
	return v.value
}

func (v *NullableCreateGroupMetaData) Set(val *CreateGroupMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupMetaData(val *CreateGroupMetaData) *NullableCreateGroupMetaData {
	return &NullableCreateGroupMetaData{value: val, isSet: true}
}

func (v NullableCreateGroupMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


