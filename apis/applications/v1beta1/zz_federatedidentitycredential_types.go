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

type FederatedIdentityCredentialObservation struct {

	// A UUID used to uniquely identify this federated identity credential
	CredentialID *string `json:"credentialId,omitempty" tf:"credential_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FederatedIdentityCredentialParameters struct {

	// The object ID of the application for which this federated identity credential should be created
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationObjectID *string `json:"applicationObjectId,omitempty" tf:"application_object_id,omitempty"`

	// Reference to a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDRef *v1.Reference `json:"applicationObjectIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDSelector *v1.Selector `json:"applicationObjectIdSelector,omitempty" tf:"-"`

	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	// +kubebuilder:validation:Required
	Audiences []*string `json:"audiences" tf:"audiences,omitempty"`

	// A description for the federated identity credential
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique display name for the federated identity credential
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	// +kubebuilder:validation:Required
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	// +kubebuilder:validation:Required
	Subject *string `json:"subject" tf:"subject,omitempty"`
}

// FederatedIdentityCredentialSpec defines the desired state of FederatedIdentityCredential
type FederatedIdentityCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FederatedIdentityCredentialParameters `json:"forProvider"`
}

// FederatedIdentityCredentialStatus defines the observed state of FederatedIdentityCredential.
type FederatedIdentityCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FederatedIdentityCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FederatedIdentityCredential is the Schema for the FederatedIdentityCredentials API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type FederatedIdentityCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FederatedIdentityCredentialSpec   `json:"spec"`
	Status            FederatedIdentityCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FederatedIdentityCredentialList contains a list of FederatedIdentityCredentials
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// Repository type metadata.
var (
	FederatedIdentityCredential_Kind             = "FederatedIdentityCredential"
	FederatedIdentityCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FederatedIdentityCredential_Kind}.String()
	FederatedIdentityCredential_KindAPIVersion   = FederatedIdentityCredential_Kind + "." + CRDGroupVersion.String()
	FederatedIdentityCredential_GroupVersionKind = CRDGroupVersion.WithKind(FederatedIdentityCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}
