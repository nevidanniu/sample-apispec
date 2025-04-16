/*
Copyright 2022.

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

package v1alpha1

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

//+genclient
//+genclient:nonNamespaced
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+k8s:openapi-gen=true

// Tenant is the Schema for Tenants Api.
// Tenant is readonly resource represents namespace, to which user has access
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TenantSpec `json:"spec,omitempty"`
}

//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TenantList contains a list of Tenants
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Tenant `json:"items"`
}

// TenantSpec defines the desired state of Tenant
type TenantSpec struct {
	corev1.NamespaceSpec `json:",inline"`
	MemberList           []TenantMember `json:"memberList,omitempty"`
}

type TenantMember struct {
	User string `json:"user"`
	Role string `json:"role"`
}

var _ resource.Object = &Tenant{}
var _ resourcestrategy.Validater = &Tenant{}

func (in *Tenant) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *Tenant) NamespaceScoped() bool {
	return false
}

func (in *Tenant) New() runtime.Object {
	return &Tenant{}
}

func (in *Tenant) NewList() runtime.Object {
	return &TenantList{}
}

func (in *Tenant) GetGroupVersionResource() schema.GroupVersionResource {
	return SchemeGroupVersion.WithResource("tenants")
}

func (in *Tenant) IsStorageVersion() bool {
	return true
}

func (in *Tenant) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &TenantList{}

func (in *TenantList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}
