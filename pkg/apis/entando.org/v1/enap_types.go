package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	GroupName string = "entando.org"
	Kind      string = "EntandoApp"
	Version   string = "v1"
	Plural    string = "entandoapps"
	Singluar  string = "entandoapp"
	ShortName string = "enap"
	Name      string = Plural + "." + GroupName
)

type EntandoAppSpec struct {
	Replicas int    `json:"replicas"`
	Dbms     string `json:"dbms"`
	// +optional
	StorageClass    string `json:"storageClass,omitempty"`
	IngressHostName string `json:"ingressHostName"`
	// +optional
	IngressPath string `json:"ingressPath,omitempty"`
	// +optional
	TlsSecretName string `json:"tlsSecretName,omitempty"`
	// +optional
	ServiceAccountToUse string `json:"serviceAccountToUse,omitempty"`
	// +optional
	EnvironmentVariables []v1.EnvVar `json:"environmentVariables,omitempty"`
	// +optional
	ResourceRequirements ResourceRequirement `json:"resourceRequirements,omitempty"`
	// +optional
	KeycloakToUse       KeycloakSpec `json:"keycloakToUse,omitempty"`
	StandardServerImage string       `json:"standardServerImage"`
	// +optional
	CustomServerImage string `json:"customServerImage"`
	// +optional
	EcrGitSshSecretName string `json:"ecrGitSshSecretName,omitempty"`
	// +optional
	BackupGitSpec BackupGitSpecification `json:"backupGitSpec,omitempty"`
}

type ResourceRequirement struct {
	StorageRequest string `json:"storageRequest"`
	StorageLimit   string `json:"storageLimit"`
	MemoryRequest  string `json:"memoryRequest"`
	MemoryLimit    string `json:"memoryLimit"`
	CpuRequest     string `json:"cpuRequest"`
	CpuLimit       string `json:"cpuLimit"`
}
type KeycloakSpec struct {
	Realm          string `json:"realm"`
	PublicClientId string `json:"publicClientId"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
}

type BackupGitSpecification struct {
	BackupGitRepo        string `json:"backupGitRepo"`
	BackupGitSecretName  string `json:"backupGitSecretName"`
	BackupResponsability string `json:"backupResponsability"`
}
type EntandoAppStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EntandoAppSpec   `json:"spec,omitempty"`
	Status EntandoAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EntandoApp `json:"items"`
}
