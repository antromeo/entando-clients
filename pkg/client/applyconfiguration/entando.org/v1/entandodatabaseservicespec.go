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

// EntandoDatabaseServiceSpecApplyConfiguration represents an declarative configuration of the EntandoDatabaseServiceSpec type for use
// with apply.
type EntandoDatabaseServiceSpecApplyConfiguration struct {
	Replicas             *int                                   `json:"replicas,omitempty"`
	Dbms                 *string                                `json:"dbms,omitempty"`
	StorageClass         *string                                `json:"storageClass,omitempty"`
	CreateDeployment     *bool                                  `json:"createDeployment,omitempty"`
	Host                 *string                                `json:"host,omitempty"`
	Port                 *int                                   `json:"port,omitempty"`
	Database             *int                                   `json:"database,omitempty"`
	TableSpace           *int                                   `json:"tableSpace,omitempty"`
	SecretName           *int                                   `json:"secretName,omitempty"`
	JdbcParameters       *int                                   `json:"jdbcParameters,omitempty"`
	ServiceAccountToUse  *int                                   `json:"serviceAccountToUse,omitempty"`
	EnvironmentVariables []v1.EnvVar                            `json:"environmentVariables,omitempty"`
	ResourceRequirements *ResourceRequirementApplyConfiguration `json:"resourceRequirements,omitempty"`
}

// EntandoDatabaseServiceSpecApplyConfiguration constructs an declarative configuration of the EntandoDatabaseServiceSpec type for use with
// apply.
func EntandoDatabaseServiceSpec() *EntandoDatabaseServiceSpecApplyConfiguration {
	return &EntandoDatabaseServiceSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithReplicas(value int) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithDbms sets the Dbms field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Dbms field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithDbms(value string) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.Dbms = &value
	return b
}

// WithStorageClass sets the StorageClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClass field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithStorageClass(value string) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.StorageClass = &value
	return b
}

// WithCreateDeployment sets the CreateDeployment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreateDeployment field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithCreateDeployment(value bool) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.CreateDeployment = &value
	return b
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithHost(value string) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.Host = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithPort(value int) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.Port = &value
	return b
}

// WithDatabase sets the Database field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Database field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithDatabase(value int) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.Database = &value
	return b
}

// WithTableSpace sets the TableSpace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TableSpace field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithTableSpace(value int) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.TableSpace = &value
	return b
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithSecretName(value int) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.SecretName = &value
	return b
}

// WithJdbcParameters sets the JdbcParameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JdbcParameters field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithJdbcParameters(value int) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.JdbcParameters = &value
	return b
}

// WithServiceAccountToUse sets the ServiceAccountToUse field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountToUse field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithServiceAccountToUse(value int) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.ServiceAccountToUse = &value
	return b
}

// WithEnvironmentVariables adds the given value to the EnvironmentVariables field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EnvironmentVariables field.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithEnvironmentVariables(values ...v1.EnvVar) *EntandoDatabaseServiceSpecApplyConfiguration {
	for i := range values {
		b.EnvironmentVariables = append(b.EnvironmentVariables, values[i])
	}
	return b
}

// WithResourceRequirements sets the ResourceRequirements field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceRequirements field is set to the value of the last call.
func (b *EntandoDatabaseServiceSpecApplyConfiguration) WithResourceRequirements(value *ResourceRequirementApplyConfiguration) *EntandoDatabaseServiceSpecApplyConfiguration {
	b.ResourceRequirements = value
	return b
}