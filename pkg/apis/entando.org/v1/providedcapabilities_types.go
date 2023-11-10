package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProvidedCapabilitySpec struct {
	Capability                string                        `json:"capability"`
	Implementation            string                        `json:"implementation"`
	Scope                     string                        `json:"scope"`
	ProvisioningStrategy      string                        `json:"provisioningStrategy"`
	Selector                  SelectorSpec                  `json:"selector"`
	CapabilityParameters      CapabilityParameters          `json:"capabilityParameters"`
	SpecifiedCapability       SpecifiedCapabilitySpec       `json:"specifiedCapability"`
	ExternallyProvidedService ExternallyProvidedServiceSpec `json:"externallyProvidedService"`
}

type SelectorSpec struct {
}

type CapabilityParameters struct {
}

type SpecifiedCapabilitySpec struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ExternallyProvidedServiceSpec struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	AdminSecretName string `json:"adminSecretName"`
}
type ProvidedCapabilityStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ProvidedCapability struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProvidedCapabilitySpec   `json:"spec,omitempty"`
	Status ProvidedCapabilityStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ProvidedCapabilityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ProvidedCapability `json:"items"`
}
