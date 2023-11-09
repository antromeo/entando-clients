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

// EntandoAppSpecApplyConfiguration represents an declarative configuration of the EntandoAppSpec type for use
// with apply.
type EntandoAppSpecApplyConfiguration struct {
	Replicas             *int                                      `json:"replicas,omitempty"`
	Dbms                 *string                                   `json:"dbms,omitempty"`
	StorageClass         *string                                   `json:"storageClass,omitempty"`
	IngressHostName      *string                                   `json:"ingressHostName,omitempty"`
	IngressPath          *string                                   `json:"ingressPath,omitempty"`
	TlsSecretName        *string                                   `json:"tlsSecretName,omitempty"`
	ServiceAccountToUse  *string                                   `json:"serviceAccountToUse,omitempty"`
	EnvironmentVariables []v1.EnvVar                               `json:"environmentVariables,omitempty"`
	ResourceRequirements *ResourceRequirementApplyConfiguration    `json:"resourceRequirements,omitempty"`
	KeycloakToUse        *KeycloakSpecApplyConfiguration           `json:"keycloakToUse,omitempty"`
	StandardServerImage  *string                                   `json:"standardServerImage,omitempty"`
	CustomServerImage    *string                                   `json:"customServerImage,omitempty"`
	EcrGitSshSecretName  *string                                   `json:"ecrGitSshSecretName,omitempty"`
	BackupGitSpec        *BackupGitSpecificationApplyConfiguration `json:"backupGitSpec,omitempty"`
}

// EntandoAppSpecApplyConfiguration constructs an declarative configuration of the EntandoAppSpec type for use with
// apply.
func EntandoAppSpec() *EntandoAppSpecApplyConfiguration {
	return &EntandoAppSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithReplicas(value int) *EntandoAppSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithDbms sets the Dbms field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Dbms field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithDbms(value string) *EntandoAppSpecApplyConfiguration {
	b.Dbms = &value
	return b
}

// WithStorageClass sets the StorageClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClass field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithStorageClass(value string) *EntandoAppSpecApplyConfiguration {
	b.StorageClass = &value
	return b
}

// WithIngressHostName sets the IngressHostName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IngressHostName field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithIngressHostName(value string) *EntandoAppSpecApplyConfiguration {
	b.IngressHostName = &value
	return b
}

// WithIngressPath sets the IngressPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IngressPath field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithIngressPath(value string) *EntandoAppSpecApplyConfiguration {
	b.IngressPath = &value
	return b
}

// WithTlsSecretName sets the TlsSecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsSecretName field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithTlsSecretName(value string) *EntandoAppSpecApplyConfiguration {
	b.TlsSecretName = &value
	return b
}

// WithServiceAccountToUse sets the ServiceAccountToUse field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountToUse field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithServiceAccountToUse(value string) *EntandoAppSpecApplyConfiguration {
	b.ServiceAccountToUse = &value
	return b
}

// WithEnvironmentVariables adds the given value to the EnvironmentVariables field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EnvironmentVariables field.
func (b *EntandoAppSpecApplyConfiguration) WithEnvironmentVariables(values ...v1.EnvVar) *EntandoAppSpecApplyConfiguration {
	for i := range values {
		b.EnvironmentVariables = append(b.EnvironmentVariables, values[i])
	}
	return b
}

// WithResourceRequirements sets the ResourceRequirements field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceRequirements field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithResourceRequirements(value *ResourceRequirementApplyConfiguration) *EntandoAppSpecApplyConfiguration {
	b.ResourceRequirements = value
	return b
}

// WithKeycloakToUse sets the KeycloakToUse field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeycloakToUse field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithKeycloakToUse(value *KeycloakSpecApplyConfiguration) *EntandoAppSpecApplyConfiguration {
	b.KeycloakToUse = value
	return b
}

// WithStandardServerImage sets the StandardServerImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StandardServerImage field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithStandardServerImage(value string) *EntandoAppSpecApplyConfiguration {
	b.StandardServerImage = &value
	return b
}

// WithCustomServerImage sets the CustomServerImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomServerImage field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithCustomServerImage(value string) *EntandoAppSpecApplyConfiguration {
	b.CustomServerImage = &value
	return b
}

// WithEcrGitSshSecretName sets the EcrGitSshSecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EcrGitSshSecretName field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithEcrGitSshSecretName(value string) *EntandoAppSpecApplyConfiguration {
	b.EcrGitSshSecretName = &value
	return b
}

// WithBackupGitSpec sets the BackupGitSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackupGitSpec field is set to the value of the last call.
func (b *EntandoAppSpecApplyConfiguration) WithBackupGitSpec(value *BackupGitSpecificationApplyConfiguration) *EntandoAppSpecApplyConfiguration {
	b.BackupGitSpec = value
	return b
}
