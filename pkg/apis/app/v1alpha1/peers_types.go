package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PeersSpec defines the desired state of Peers
// +k8s:openapi-gen=true
type PeersSpec struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   peers_metadata `json:"metadata"`
	Replicas string `json:"replicas"`
	Spec peers_specs `json:"spec"`
}

type peers_metadata struct {
		Name   string `json:"name"`
		Labels struct {
			App    string `json:"app"`
			Role   string `json:"role"`
			PeerId string `json:"peerId"`
			Org    string `json:"org"`
		} `json:"labels"`
	}

type peers_specs struct {
		RestartPolicy string `json:"restartPolicy"`
		Containers    []struct {
			Name            string `json:"name"`
			ImagePullPolicy string `json:"imagePullPolicy"`
			Image           string `json:"image"`
			WorkingDir      string `json:"workingDir"`
			VolumeMounts    []struct {
				MountPath string `json:"mountPath"`
				Name      string `json:"name"`
			} `json:"volumeMounts"`
			Env []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"env"`
			Ports []struct {
				ContainerPort int `json:"containerPort"`
			} `json:"ports"`
			Command []string `json:"command"`
		} `json:"containers"`
		Volumes []struct {
			Name     string `json:"name"`
			HostPath struct {
				Path string `json:"path"`
			} `json:"hostPath"`
		} `json:"volumes"`
	}

// PeersStatus defines the observed state of Peers
// +k8s:openapi-gen=true
type PeersStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Peers []string `json:"peers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Peers is the Schema for the peers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Peers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PeersSpec   `json:"spec,omitempty"`
	Status PeersStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PeersList contains a list of Peers
type PeersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Peers `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Peers{}, &PeersList{})
}
