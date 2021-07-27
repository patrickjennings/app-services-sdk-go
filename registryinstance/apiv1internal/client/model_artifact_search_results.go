/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.1.0-SNAPSHOT
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// ArtifactSearchResults Describes the response received when searching for artifacts.
type ArtifactSearchResults struct {

	// The artifacts returned in the result set.
	Artifacts []SearchedArtifact `json:"artifacts"`

	// The total number of artifacts that matched the query that produced the result set (may be  more than the number of artifacts in the result set).
	Count int32 `json:"count"`

}

// NewArtifactSearchResults instantiates a new ArtifactSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactSearchResults(artifacts []SearchedArtifact, count int32) *ArtifactSearchResults {
	this := ArtifactSearchResults{}
	this.Artifacts = artifacts
	this.Count = count
	return &this
}

// NewArtifactSearchResultsWithDefaults instantiates a new ArtifactSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactSearchResultsWithDefaults() *ArtifactSearchResults {
	this := ArtifactSearchResults{}



	return &this
}


// GetArtifacts returns the Artifacts field value
func (o *ArtifactSearchResults) GetArtifacts() []SearchedArtifact {
	if o == nil {
		var ret []SearchedArtifact
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *ArtifactSearchResults) GetArtifactsOk() (*[]SearchedArtifact, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Artifacts, true
}

// SetArtifacts sets field value
func (o *ArtifactSearchResults) SetArtifacts(v []SearchedArtifact) {
	o.Artifacts = v
}


// GetCount returns the Count field value
func (o *ArtifactSearchResults) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ArtifactSearchResults) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ArtifactSearchResults) SetCount(v int32) {
	o.Count = v
}


func (o ArtifactSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["artifacts"] = o.Artifacts
	}
    
	if true {
		toSerialize["count"] = o.Count
	}
    
	return json.Marshal(toSerialize)
}

type NullableArtifactSearchResults struct {
	value *ArtifactSearchResults
	isSet bool
}

func (v NullableArtifactSearchResults) Get() *ArtifactSearchResults {
	return v.value
}

func (v *NullableArtifactSearchResults) Set(val *ArtifactSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactSearchResults(val *ArtifactSearchResults) *NullableArtifactSearchResults {
	return &NullableArtifactSearchResults{value: val, isSet: true}
}

func (v NullableArtifactSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

