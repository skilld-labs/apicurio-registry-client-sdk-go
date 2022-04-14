/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.3.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
)

// ContentCreateRequest struct for ContentCreateRequest
type ContentCreateRequest struct {
	// Raw content of the artifact.
	Content string `json:"content"`
	// Collection of references to other artifacts.
	References []ArtifactReference `json:"references"`
}

// NewContentCreateRequest instantiates a new ContentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentCreateRequest(content string, references []ArtifactReference) *ContentCreateRequest {
	this := ContentCreateRequest{}
	this.Content = content
	this.References = references
	return &this
}

// NewContentCreateRequestWithDefaults instantiates a new ContentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentCreateRequestWithDefaults() *ContentCreateRequest {
	this := ContentCreateRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *ContentCreateRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ContentCreateRequest) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ContentCreateRequest) SetContent(v string) {
	o.Content = v
}

// GetReferences returns the References field value
func (o *ContentCreateRequest) GetReferences() []ArtifactReference {
	if o == nil {
		var ret []ArtifactReference
		return ret
	}

	return o.References
}

// GetReferencesOk returns a tuple with the References field value
// and a boolean to check if the value has been set.
func (o *ContentCreateRequest) GetReferencesOk() ([]ArtifactReference, bool) {
	if o == nil  {
		return nil, false
	}
	return o.References, true
}

// SetReferences sets field value
func (o *ContentCreateRequest) SetReferences(v []ArtifactReference) {
	o.References = v
}

func (o ContentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableContentCreateRequest struct {
	value *ContentCreateRequest
	isSet bool
}

func (v NullableContentCreateRequest) Get() *ContentCreateRequest {
	return v.value
}

func (v *NullableContentCreateRequest) Set(val *ContentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentCreateRequest(val *ContentCreateRequest) *NullableContentCreateRequest {
	return &NullableContentCreateRequest{value: val, isSet: true}
}

func (v NullableContentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


