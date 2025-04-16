/*
Copyright 2021.

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

package v1alpha1

// TenantSpecApplyConfiguration represents an declarative configuration of the TenantSpec type for use
// with apply.
type TenantSpecApplyConfiguration struct {
	unnameable_Unsupported `json:",inline"`
	MemberList             []TenantMemberApplyConfiguration `json:"memberList,omitempty"`
}

// TenantSpecApplyConfiguration constructs an declarative configuration of the TenantSpec type for use with
// apply.
func TenantSpec() *TenantSpecApplyConfiguration {
	return &TenantSpecApplyConfiguration{}
}

// WithMemberList adds the given value to the MemberList field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MemberList field.
func (b *TenantSpecApplyConfiguration) WithMemberList(values ...*TenantMemberApplyConfiguration) *TenantSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMemberList")
		}
		b.MemberList = append(b.MemberList, *values[i])
	}
	return b
}
