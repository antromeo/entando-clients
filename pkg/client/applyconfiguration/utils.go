/*
Copyright The Kubernetes Authors.

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

package applyconfiguration

import (
	v1 "github.com/antromeo/entando-clients/pkg/apis/entando.org/v1"
	entandoorgv1 "github.com/antromeo/entando-clients/pkg/client/applyconfiguration/entando.org/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=entando.org, Version=v1
	case v1.SchemeGroupVersion.WithKind("BackupGitSpecification"):
		return &entandoorgv1.BackupGitSpecificationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DetailSpec"):
		return &entandoorgv1.DetailSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DistTag"):
		return &entandoorgv1.DistTagApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoApp"):
		return &entandoorgv1.EntandoAppApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoAppPluginLink"):
		return &entandoorgv1.EntandoAppPluginLinkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoAppPluginLinkSpec"):
		return &entandoorgv1.EntandoAppPluginLinkSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoAppPluginLinkStatus"):
		return &entandoorgv1.EntandoAppPluginLinkStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoAppSpec"):
		return &entandoorgv1.EntandoAppSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoAppStatus"):
		return &entandoorgv1.EntandoAppStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoDatabaseService"):
		return &entandoorgv1.EntandoDatabaseServiceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoDatabaseServiceSpec"):
		return &entandoorgv1.EntandoDatabaseServiceSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoDatabaseServiceStatus"):
		return &entandoorgv1.EntandoDatabaseServiceStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoDeBundle"):
		return &entandoorgv1.EntandoDeBundleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoDeBundleSpec"):
		return &entandoorgv1.EntandoDeBundleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoDeBundleStatus"):
		return &entandoorgv1.EntandoDeBundleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoKeycloakServer"):
		return &entandoorgv1.EntandoKeycloakServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoKeycloakServerSpec"):
		return &entandoorgv1.EntandoKeycloakServerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoKeycloakServerStatus"):
		return &entandoorgv1.EntandoKeycloakServerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoPlugin"):
		return &entandoorgv1.EntandoPluginApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoPluginSpec"):
		return &entandoorgv1.EntandoPluginSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EntandoPluginStatus"):
		return &entandoorgv1.EntandoPluginStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternallyProvidedServiceSpec"):
		return &entandoorgv1.ExternallyProvidedServiceSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KeycloakSpec"):
		return &entandoorgv1.KeycloakSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Permission"):
		return &entandoorgv1.PermissionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProvidedCapability"):
		return &entandoorgv1.ProvidedCapabilityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProvidedCapabilitySpec"):
		return &entandoorgv1.ProvidedCapabilitySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProvidedCapabilityStatus"):
		return &entandoorgv1.ProvidedCapabilityStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ResourceRequirement"):
		return &entandoorgv1.ResourceRequirementApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Role"):
		return &entandoorgv1.RoleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SpecifiedCapabilitySpec"):
		return &entandoorgv1.SpecifiedCapabilitySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TagsSpec"):
		return &entandoorgv1.TagsSpecApplyConfiguration{}

	}
	return nil
}
