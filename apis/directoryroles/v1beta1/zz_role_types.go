/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RoleObservation struct {

	// The description of the directory role
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of the directory role
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The object ID of the template associated with the directory role
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type RoleParameters struct {

	// The display name of the directory role
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Role is the Schema for the Roles API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleSpec   `json:"spec"`
	Status            RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
