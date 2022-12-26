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

type RoleAssignmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleAssignmentParameters struct {

	// Identifier of the app-specific scope when the assignment scope is app-specific
	// +kubebuilder:validation:Optional
	AppScopeID *string `json:"appScopeId,omitempty" tf:"app_scope_id,omitempty"`

	// Identifier of the app-specific scope when the assignment scope is app-specific
	// +kubebuilder:validation:Optional
	AppScopeObjectID *string `json:"appScopeObjectId,omitempty" tf:"app_scope_object_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment
	// +kubebuilder:validation:Optional
	DirectoryScopeID *string `json:"directoryScopeId,omitempty" tf:"directory_scope_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment
	// +kubebuilder:validation:Optional
	DirectoryScopeObjectID *string `json:"directoryScopeObjectId,omitempty" tf:"directory_scope_object_id,omitempty"`

	// The object ID of the member principal
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/users/v1beta1.User
	// +kubebuilder:validation:Optional
	PrincipalObjectID *string `json:"principalObjectId,omitempty" tf:"principal_object_id,omitempty"`

	// Reference to a User in users to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDRef *v1.Reference `json:"principalObjectIdRef,omitempty" tf:"-"`

	// Selector for a User in users to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDSelector *v1.Selector `json:"principalObjectIdSelector,omitempty" tf:"-"`

	// The object ID of the directory role for this assignment
	// +crossplane:generate:reference:type=Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("template_id",true)
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Reference to a Role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDRef *v1.Reference `json:"roleIdRef,omitempty" tf:"-"`

	// Selector for a Role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`
}

// RoleAssignmentSpec defines the desired state of RoleAssignment
type RoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAssignmentParameters `json:"forProvider"`
}

// RoleAssignmentStatus defines the observed state of RoleAssignment.
type RoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignment is the Schema for the RoleAssignments API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignmentSpec   `json:"spec"`
	Status            RoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentList contains a list of RoleAssignments
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	RoleAssignment_Kind             = "RoleAssignment"
	RoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAssignment_Kind}.String()
	RoleAssignment_KindAPIVersion   = RoleAssignment_Kind + "." + CRDGroupVersion.String()
	RoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(RoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
