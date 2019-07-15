package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CAsSpec defines the desired state of CAs
// +k8s:openapi-gen=true
type CAsSpec struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name   string `json:"name"`
		Labels struct {
			App  string `json:"app"`
			Role string `json:"role"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Replicas int `json:"replicas"`
		Template struct {
			Metadata struct {
				Name   string `json:"name"`
				Labels struct {
					Role string `json:"role"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				RestartPolicy string `json:"restartPolicy"`
				Containers    []struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					Ports []struct {
						ContainerPort int `json:"containerPort"`
					} `json:"ports"`
					Command []string `json:"command"`
				} `json:"containers"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
}

// CAsStatus defines the observed state of CAs
// +k8s:openapi-gen=true
type CAsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	CAs []string `json:"cas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CAs is the Schema for the cas API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type CAs struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CAsSpec   `json:"spec,omitempty"`
	Status CAsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CAsList contains a list of CAs
type CAsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CAs `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CAs{}, &CAsList{})
}
