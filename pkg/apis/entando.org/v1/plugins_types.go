package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EntandoPluginSpec struct {
	// +optional
	TenantCode          string `json:"tenantCode,omitempty"`
	Image               string `json:"spec"`
	Replicas            int    `json:"replicas"`
	Dbms                string `json:"dbms"`
	StorageClass        string `json:"storageClass"`
	IngressHostName     string `json:"ingressHostName"`
	IngressPath         string `json:"ingressPath"`
	TlsSecretName       string `json:"tlsSecretName"`
	ServiceAccountToUse string `json:"serviceAccountToUse"`
	// +optional
	EnvironmentVariables []v1.EnvVar `json:"environmentVariables,omitempty"`
	// +optional
	ResourceRequirements ResourceRequirement `json:"resourceRequirements,omitempty"`
	// +optional
	KeycloakToUse        KeycloakSpec `json:"keycloakToUse,omitempty"`
	HealthCheckPath      string       `json:"healthCheckPath"`
	SecurityLevel        string       `json:"securityLevel"`
	ConnectionConfigName string       `json:"connectionConfigNames"`
	CompanionContainers  []string     `json:"companionContainers"`
	Roles                []Role       `json:"roles"`
	Permissions          []Permission `json:"permissions"`
}

type Role struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Permission struct {
	ClientId string `json:"clientId"`
	Role     string `json:"role"`
}

type EntandoPluginStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EntandoPluginSpec   `json:"spec,omitempty"`
	Status EntandoPluginStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EntandoPlugin `json:"items"`
}
