package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EntandoDeBundleSpec struct {
	Details DetailSpec `json:"details"`
	Tags    []TagsSpec `json:"spec"`
}

type DetailSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	DistTags DistTag  `json:"dist-tags"`
	Versions []string `json:"versions"`
	// +optional
	Thumbnail string `json:"thumbnail,omitempty"`
}

type DistTag struct {
	Latest string `json:"latest"`
}

type TagsSpec struct {
	Version   string `json:"version"`
	Integrity string `json:"integrity"`
	Shasum    string `json:"shasum"`
	Tarball   string `json:"tarball"`
}

type EntandoDeBundleStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoDeBundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EntandoDeBundleSpec   `json:"spec,omitempty"`
	Status EntandoDeBundleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoDeBundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EntandoDeBundle `json:"items"`
}
