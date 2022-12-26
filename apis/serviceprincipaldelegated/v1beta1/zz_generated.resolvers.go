/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this PermissionGrant.
func (mg *PermissionGrant) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceServicePrincipalObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceServicePrincipalObjectIDRef,
		Selector:     mg.Spec.ForProvider.ResourceServicePrincipalObjectIDSelector,
		To: reference.To{
			List:    &v1beta1.PrincipalList{},
			Managed: &v1beta1.Principal{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceServicePrincipalObjectID")
	}
	mg.Spec.ForProvider.ResourceServicePrincipalObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceServicePrincipalObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServicePrincipalObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServicePrincipalObjectIDRef,
		Selector:     mg.Spec.ForProvider.ServicePrincipalObjectIDSelector,
		To: reference.To{
			List:    &v1beta1.PrincipalList{},
			Managed: &v1beta1.Principal{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServicePrincipalObjectID")
	}
	mg.Spec.ForProvider.ServicePrincipalObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServicePrincipalObjectIDRef = rsp.ResolvedReference

	return nil
}
