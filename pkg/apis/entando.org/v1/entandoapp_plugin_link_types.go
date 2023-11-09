package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EntandoAppPluginLinkSpec struct {
	EntandoAppName         string `json:"entandoAppName"`
	EntandoAppNamespace    string `json:"entandoAppNamespace"`
	EntandoPluginName      string `json:"entandoPluginName"`
	EntandoPluginNamespace string `json:"entandoPluginNamespace"`
	// +optional
	TenantCode string `json:"tenantCode,omitempty"`
}

type EntandoAppPluginLinkStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoAppPluginLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EntandoAppPluginLinkSpec   `json:"spec,omitempty"`
	Status EntandoAppPluginLinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EntandoAppPluginLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EntandoAppPluginLink `json:"items"`
}
