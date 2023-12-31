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

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// EntandoKeycloakServerSpecApplyConfiguration represents an declarative configuration of the EntandoKeycloakServerSpec type for use
// with apply.
type EntandoKeycloakServerSpecApplyConfiguration struct {
	Replicas                *int                                   `json:"replicas,omitempty"`
	Dbms                    *string                                `json:"dbms,omitempty"`
	StorageClass            *string                                `json:"storageClass,omitempty"`
	IngressHostName         *string                                `json:"ingressHostName,omitempty"`
	TlsSecretName           *string                                `json:"tlsSecretName,omitempty"`
	ServiceAccountToUse     *string                                `json:"serviceAccountToUse,omitempty"`
	EnvironmentVariables    []v1.EnvVar                            `json:"environmentVariables,omitempty"`
	ResourceRequirements    *ResourceRequirementApplyConfiguration `json:"resourceRequirements,omitempty"`
	CustomImage             *string                                `json:"customImage,omitempty"`
	FrontEndUrl             *string                                `json:"frontEndUrl,omitempty"`
	IsDefault               *bool                                  `json:"isDefault,omitempty"`
	StandardImage           *string                                `json:"standardImage,omitempty"`
	ProvidedCapabilityScope *string                                `json:"providedCapabilityScope,omitempty"`
	ProvisioningStrategy    *string                                `json:"provisioningStrategy,omitempty"`
	AdminSecretName         *string                                `json:"adminSecretName,omitempty"`
	DefaultRealm            *string                                `json:"defaultRealm,omitempty"`
}

// EntandoKeycloakServerSpecApplyConfiguration constructs an declarative configuration of the EntandoKeycloakServerSpec type for use with
// apply.
func EntandoKeycloakServerSpec() *EntandoKeycloakServerSpecApplyConfiguration {
	return &EntandoKeycloakServerSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithReplicas(value int) *EntandoKeycloakServerSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithDbms sets the Dbms field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Dbms field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithDbms(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.Dbms = &value
	return b
}

// WithStorageClass sets the StorageClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClass field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithStorageClass(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.StorageClass = &value
	return b
}

// WithIngressHostName sets the IngressHostName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IngressHostName field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithIngressHostName(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.IngressHostName = &value
	return b
}

// WithTlsSecretName sets the TlsSecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsSecretName field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithTlsSecretName(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.TlsSecretName = &value
	return b
}

// WithServiceAccountToUse sets the ServiceAccountToUse field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountToUse field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithServiceAccountToUse(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.ServiceAccountToUse = &value
	return b
}

// WithEnvironmentVariables adds the given value to the EnvironmentVariables field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EnvironmentVariables field.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithEnvironmentVariables(values ...v1.EnvVar) *EntandoKeycloakServerSpecApplyConfiguration {
	for i := range values {
		b.EnvironmentVariables = append(b.EnvironmentVariables, values[i])
	}
	return b
}

// WithResourceRequirements sets the ResourceRequirements field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceRequirements field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithResourceRequirements(value *ResourceRequirementApplyConfiguration) *EntandoKeycloakServerSpecApplyConfiguration {
	b.ResourceRequirements = value
	return b
}

// WithCustomImage sets the CustomImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomImage field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithCustomImage(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.CustomImage = &value
	return b
}

// WithFrontEndUrl sets the FrontEndUrl field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FrontEndUrl field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithFrontEndUrl(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.FrontEndUrl = &value
	return b
}

// WithIsDefault sets the IsDefault field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IsDefault field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithIsDefault(value bool) *EntandoKeycloakServerSpecApplyConfiguration {
	b.IsDefault = &value
	return b
}

// WithStandardImage sets the StandardImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StandardImage field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithStandardImage(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.StandardImage = &value
	return b
}

// WithProvidedCapabilityScope sets the ProvidedCapabilityScope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProvidedCapabilityScope field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithProvidedCapabilityScope(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.ProvidedCapabilityScope = &value
	return b
}

// WithProvisioningStrategy sets the ProvisioningStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProvisioningStrategy field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithProvisioningStrategy(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.ProvisioningStrategy = &value
	return b
}

// WithAdminSecretName sets the AdminSecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminSecretName field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithAdminSecretName(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.AdminSecretName = &value
	return b
}

// WithDefaultRealm sets the DefaultRealm field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultRealm field is set to the value of the last call.
func (b *EntandoKeycloakServerSpecApplyConfiguration) WithDefaultRealm(value string) *EntandoKeycloakServerSpecApplyConfiguration {
	b.DefaultRealm = &value
	return b
}
