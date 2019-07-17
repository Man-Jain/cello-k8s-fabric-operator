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
	Metadata metadata `json:"metadata"`
	Spec spec `json:"spec"`
}

type metadata struct {
		Name   string `json:"name"`
		Labels struct {
			App    string `json:"app"`
			Role   string `json:"role"`
			PeerID string `json:"peer-id"`
			Org    string `json:"org"`
		} `json:"labels"`
}

type Volumemounts struct {
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
}

type Envs struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CPorts struct {
	ContainerPort int `json:"containerPort"`
}

type CVolumes struct {
	Name     string `json:"name"`
	HostPath struct {
		Path string `json:"path"`
	} `json:"hostPath"`
}

type CContainers struct {
					Name            string `json:"name"`
					ImagePullPolicy string `json:"imagePullPolicy"`
					Image           string `json:"image"`
					WorkingDir      string `json:"workingDir"`
					VolumeMounts []Volumemounts `json:"volumeMounts"`
					Env []Envs `json:"env"`
					Ports []CPorts `json:"ports"`
					Command []string `json:"command"`
				}

type spec struct {
		Replicas int `json:"replicas"`
		Template struct {
			Metadata struct {
				Name   string `json:"name"`
				Labels struct {
					App    string `json:"app"`
					Role   string `json:"role"`
					PeerID string `json:"peer-id"`
					Org    string `json:"org"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				RestartPolicy string `json:"restartPolicy"`
				Containers CContainers `json:"containers"`
				Volumes []CVolumes `json:"volumes"`
			} `json:"spec"`
		} `json:"template"`
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
