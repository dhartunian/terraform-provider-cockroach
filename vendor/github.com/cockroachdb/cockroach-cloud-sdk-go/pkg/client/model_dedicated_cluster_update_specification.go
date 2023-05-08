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

// DedicatedClusterUpdateSpecification struct for DedicatedClusterUpdateSpecification.
type DedicatedClusterUpdateSpecification struct {
	// This field should contain the CMEK specs for newly added regions. If a CMEK spec is provided for an existing region, the request is invalid and will fail.
	CmekRegionSpecs *[]CMEKRegionSpecification            `json:"cmek_region_specs,omitempty"`
	Hardware        *DedicatedHardwareUpdateSpecification `json:"hardware,omitempty"`
	// Region keys should match the cloud provider's zone code. For example, for Oregon, set region_name to \"us-west2\" for GCP and \"us-west-2\" for AWS. Values represent the node count.
	RegionNodes *map[string]int32 `json:"region_nodes,omitempty"`
}

// NewDedicatedClusterUpdateSpecification instantiates a new DedicatedClusterUpdateSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedClusterUpdateSpecification() *DedicatedClusterUpdateSpecification {
	p := DedicatedClusterUpdateSpecification{}
	return &p
}

// GetCmekRegionSpecs returns the CmekRegionSpecs field value if set, zero value otherwise.
func (o *DedicatedClusterUpdateSpecification) GetCmekRegionSpecs() []CMEKRegionSpecification {
	if o == nil || o.CmekRegionSpecs == nil {
		var ret []CMEKRegionSpecification
		return ret
	}
	return *o.CmekRegionSpecs
}

// SetCmekRegionSpecs gets a reference to the given []CMEKRegionSpecification and assigns it to the CmekRegionSpecs field.
func (o *DedicatedClusterUpdateSpecification) SetCmekRegionSpecs(v []CMEKRegionSpecification) {
	o.CmekRegionSpecs = &v
}

// GetHardware returns the Hardware field value if set, zero value otherwise.
func (o *DedicatedClusterUpdateSpecification) GetHardware() DedicatedHardwareUpdateSpecification {
	if o == nil || o.Hardware == nil {
		var ret DedicatedHardwareUpdateSpecification
		return ret
	}
	return *o.Hardware
}

// SetHardware gets a reference to the given DedicatedHardwareUpdateSpecification and assigns it to the Hardware field.
func (o *DedicatedClusterUpdateSpecification) SetHardware(v DedicatedHardwareUpdateSpecification) {
	o.Hardware = &v
}

// GetRegionNodes returns the RegionNodes field value if set, zero value otherwise.
func (o *DedicatedClusterUpdateSpecification) GetRegionNodes() map[string]int32 {
	if o == nil || o.RegionNodes == nil {
		var ret map[string]int32
		return ret
	}
	return *o.RegionNodes
}

// SetRegionNodes gets a reference to the given map[string]int32 and assigns it to the RegionNodes field.
func (o *DedicatedClusterUpdateSpecification) SetRegionNodes(v map[string]int32) {
	o.RegionNodes = &v
}
