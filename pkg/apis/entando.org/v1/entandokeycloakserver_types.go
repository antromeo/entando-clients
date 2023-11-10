package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EntandoKeycloakServerSpec struct {
	Replicas            int    `json:"replicas"`
	Dbms                string `json:"dbms"`
	StorageClass        string `json:"storageClass"`
	IngressHostName     string `json:"ingressHostName"`
	TlsSecretName       string `json:"tlsSecretName"`
	ServiceAccountToUse string `json:"serviceAccountToUse"`
	// +optional
	EnvironmentVariables []v1.EnvVar `json:"environmentVariables,omitempty"`
	// +optional
	ResourceRequirements    ResourceRequirement `json:"resourceRequirements,omitempty"`
	CustomImage             string              `json:"customImage"`
	FrontEndUrl             string              `json:"frontEndUrl"`
	IsDefault               bool                `json:"isDefault"`
	StandardImage           string              `json:"standardImage"`
	ProvidedCapabilityScope string              `json:"providedCapabilityScope"`
	ProvisioningStrategy    string              `json:"provisioningStrategy"`
	AdminSecretName         string              `json:"adminSecretName"`
	DefaultRealm            string              `json:"defaultRealm"`
}

type EntandoKeycloakServerStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoKeycloakServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EntandoKeycloakServerSpec   `json:"spec,omitempty"`
	Status EntandoKeycloakServerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoKeycloakServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EntandoKeycloakServer `json:"items"`
}
