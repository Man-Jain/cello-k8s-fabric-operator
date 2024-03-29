package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OrderersSpec defines the desired state of Orderers
// +k8s:openapi-gen=true
type OrderersSpec struct {
	Metadata   *OrderersMetadata `json:"metadata"`
	RestartPolicy string `json:"restartPolicy"`
	Containers    []*OrderersContainers `json:"containers"`
	Volumes []*OrderersVolumes `json:"volumes"`
}

type OrderersMetadata struct {
	Name   string `json:"name"`
	Labels *OrderersLabels `json:"labels"`
}

type OrderersLabels struct {
	App    string `json:"app"`
	Role   string `json:"role"`
}

type OrderersContainers struct {
	Name            string `json:"name"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	Image           string `json:"image"`
	WorkingDir      string `json:"workingDir"`
	VolumeMounts    []*OrderersVolumeMounts `json:"volumeMounts"`
	Env []*OrderersEnv `json:"env"`
	Ports []*OrderersPorts `json:"ports"`
	Command []string `json:"command"`
}

type OrderersVolumeMounts struct {
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
}

type OrderersEnv struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OrderersPorts struct {
	ContainerPort int `json:"containerPort"`
}

type OrderersVolumes struct {
	Name     string `json:"name"`
	HostPath *OrderersHostpath `json:"hostPath"`
}

type OrderersHostpath struct {
	Path string `json:"path"`
}

// OrderersStatus defines the observed state of Orderers
// +k8s:openapi-gen=true
type OrderersStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Orderers []string `json:"orderers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Orderers is the Schema for the orderers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Orderers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrderersSpec   `json:"spec,omitempty"`
	Status OrderersStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrderersList contains a list of Orderers
type OrderersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Orderers `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Orderers{}, &OrderersList{})
}
