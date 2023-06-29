/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
)

// ResourceFlavorSpecApplyConfiguration represents an declarative configuration of the ResourceFlavorSpec type for use
// with apply.
type ResourceFlavorSpecApplyConfiguration struct {
	NodeLabels map[string]string `json:"nodeLabels,omitempty"`
	NodeTaints []v1.Taint        `json:"nodeTaints,omitempty"`
}

// ResourceFlavorSpecApplyConfiguration constructs an declarative configuration of the ResourceFlavorSpec type for use with
// apply.
func ResourceFlavorSpec() *ResourceFlavorSpecApplyConfiguration {
	return &ResourceFlavorSpecApplyConfiguration{}
}

// WithNodeLabels puts the entries into the NodeLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeLabels field,
// overwriting an existing map entries in NodeLabels field with the same key.
func (b *ResourceFlavorSpecApplyConfiguration) WithNodeLabels(entries map[string]string) *ResourceFlavorSpecApplyConfiguration {
	if b.NodeLabels == nil && len(entries) > 0 {
		b.NodeLabels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeLabels[k] = v
	}
	return b
}

// WithNodeTaints adds the given value to the NodeTaints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NodeTaints field.
func (b *ResourceFlavorSpecApplyConfiguration) WithNodeTaints(values ...v1.Taint) *ResourceFlavorSpecApplyConfiguration {
	for i := range values {
		b.NodeTaints = append(b.NodeTaints, values[i])
	}
	return b
}