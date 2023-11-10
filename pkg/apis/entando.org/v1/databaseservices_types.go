package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EntandoDatabaseServiceSpec struct {
	Replicas            int            `json:"replicas"`
	Dbms                string         `json:"dbms"`
	StorageClass        string         `json:"storageClass"`
	CreateDeployment    bool           `json:"createDeployment"`
	Host                string         `json:"host"`
	Port                int            `json:"port"`
	Database            string         `json:"database"`
	TableSpace          string         `json:"tableSpace"`
	SecretName          string         `json:"secretName"`
	JdbcParameters      JdbcParameters `json:"jdbcParameters"`
	ServiceAccountToUse string         `json:"serviceAccountToUse"`
	// +optional
	EnvironmentVariables []v1.EnvVar `json:"environmentVariables,omitempty"`
	// +optional
	ResourceRequirements ResourceRequirement `json:"resourceRequirements,omitempty"`
}

type JdbcParameters interface {
}

type EntandoDatabaseServiceStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoDatabaseService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EntandoDatabaseServiceSpec   `json:"spec,omitempty"`
	Status EntandoDatabaseServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoDatabaseServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EntandoDatabaseService `json:"items"`
}
