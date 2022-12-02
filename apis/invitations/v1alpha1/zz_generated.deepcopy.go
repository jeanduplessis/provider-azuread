//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Invitation) DeepCopyInto(out *Invitation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Invitation.
func (in *Invitation) DeepCopy() *Invitation {
	if in == nil {
		return nil
	}
	out := new(Invitation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Invitation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvitationList) DeepCopyInto(out *InvitationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Invitation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvitationList.
func (in *InvitationList) DeepCopy() *InvitationList {
	if in == nil {
		return nil
	}
	out := new(InvitationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InvitationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvitationObservation) DeepCopyInto(out *InvitationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RedeemURL != nil {
		in, out := &in.RedeemURL, &out.RedeemURL
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvitationObservation.
func (in *InvitationObservation) DeepCopy() *InvitationObservation {
	if in == nil {
		return nil
	}
	out := new(InvitationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvitationParameters) DeepCopyInto(out *InvitationParameters) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = make([]MessageParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RedirectURL != nil {
		in, out := &in.RedirectURL, &out.RedirectURL
		*out = new(string)
		**out = **in
	}
	if in.UserDisplayName != nil {
		in, out := &in.UserDisplayName, &out.UserDisplayName
		*out = new(string)
		**out = **in
	}
	if in.UserEmailAddress != nil {
		in, out := &in.UserEmailAddress, &out.UserEmailAddress
		*out = new(string)
		**out = **in
	}
	if in.UserType != nil {
		in, out := &in.UserType, &out.UserType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvitationParameters.
func (in *InvitationParameters) DeepCopy() *InvitationParameters {
	if in == nil {
		return nil
	}
	out := new(InvitationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvitationSpec) DeepCopyInto(out *InvitationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvitationSpec.
func (in *InvitationSpec) DeepCopy() *InvitationSpec {
	if in == nil {
		return nil
	}
	out := new(InvitationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvitationStatus) DeepCopyInto(out *InvitationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvitationStatus.
func (in *InvitationStatus) DeepCopy() *InvitationStatus {
	if in == nil {
		return nil
	}
	out := new(InvitationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessageObservation) DeepCopyInto(out *MessageObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessageObservation.
func (in *MessageObservation) DeepCopy() *MessageObservation {
	if in == nil {
		return nil
	}
	out := new(MessageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessageParameters) DeepCopyInto(out *MessageParameters) {
	*out = *in
	if in.AdditionalRecipients != nil {
		in, out := &in.AdditionalRecipients, &out.AdditionalRecipients
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessageParameters.
func (in *MessageParameters) DeepCopy() *MessageParameters {
	if in == nil {
		return nil
	}
	out := new(MessageParameters)
	in.DeepCopyInto(out)
	return out
}