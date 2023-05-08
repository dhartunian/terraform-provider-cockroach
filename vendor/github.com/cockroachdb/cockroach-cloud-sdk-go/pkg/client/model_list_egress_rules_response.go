// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2023-04-10

package client

// ListEgressRulesResponse ListEgressRulesResponse is the output for the ListEgressRules RPC..
type ListEgressRulesResponse struct {
	Pagination *KeysetPaginationResponse `json:"pagination,omitempty"`
	// rules are the egress rules associated with the given CockroachDB cluster.
	Rules *[]EgressRule `json:"rules,omitempty"`
}

// NewListEgressRulesResponse instantiates a new ListEgressRulesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEgressRulesResponse() *ListEgressRulesResponse {
	p := ListEgressRulesResponse{}
	return &p
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListEgressRulesResponse) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ListEgressRulesResponse) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ListEgressRulesResponse) GetRules() []EgressRule {
	if o == nil || o.Rules == nil {
		var ret []EgressRule
		return ret
	}
	return *o.Rules
}

// SetRules gets a reference to the given []EgressRule and assigns it to the Rules field.
func (o *ListEgressRulesResponse) SetRules(v []EgressRule) {
	o.Rules = &v
}
