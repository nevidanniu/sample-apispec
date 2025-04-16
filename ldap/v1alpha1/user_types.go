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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
)

//+genclient
//+genclient:nonNamespaced
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+k8s:openapi-gen=true

// User is the Schema for Users Api.
// User is readonly resource represents namespace, to which user has access
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec       UserSpec   `json:"spec,omitempty"`
	UserStatus UserStatus `json:"status,omitempty"`
}

//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []User `json:"items"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	// CN - имя пользователя
	CN string `json:"cn"`
	// SN - фамилия пользователя
	SN string `json:"sn"`
	// Mail - почта пользователя
	Mail string `json:"mail"`
	// Uid Username (логин)
	Uid string `json:"uid"`
}

// UserStatus defines the observed state of User
type UserStatus struct {
	// DN Distinguished Name
	DN string `json:"dn,omitempty"`
}

var _ resource.Object = &User{}

func (in *User) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *User) NamespaceScoped() bool {
	return false
}

func (in *User) New() runtime.Object {
	return &User{}
}

func (in *User) NewList() runtime.Object {
	return &UserList{}
}

func (in *User) GetGroupVersionResource() schema.GroupVersionResource {
	return SchemeGroupVersion.WithResource("users")
}

func (in *User) IsStorageVersion() bool {
	return true
}

var _ resource.ObjectList = &UserList{}

func (in *UserList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}
