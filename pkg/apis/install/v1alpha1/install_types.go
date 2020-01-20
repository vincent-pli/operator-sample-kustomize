package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InstallSpec defines the desired state of Install
type InstallSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	TargetNamespace string    `json:"targetNamespace" protobuf:"bytes,3,opt,name=targetNamespace"`
	SetOwner        bool      `json:"setowner" protobuf:"bytes,3,opt,name=setowner"`
	Registry        *Registry `json:"registry,omitempty"`
}

// Registry defines image overrides of knative images.
// The default value is used as a default format to override for all knative deployments.
// The override values are specific to each knative deployment.
// +k8s:openapi-gen=true
type Registry struct {
	// A map of a container name or arg key to the full image location of the individual knative container.
	// +optional
	Override map[string]string `json:"override,omitempty"`
}

type InstallState string

const (
	InstallStateError      InstallState = "Error"
	InstallStateInstalling InstallState = "Installing"
	InstallStateInstalled  InstallState = "Installed"
)

// InstallStatus defines the observed state of Install
type InstallStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	State   InstallState `json:"state,omitempty"`
	Version string       `json:"version,omitempty"`
	Message string       `json:"message,omitempty"`
}

func (is *InstallStatus) MarkInstallFailed(msg string) {
	is.State = InstallStateError
	is.Message = msg
	is.Version = "-"
}

func (is *InstallStatus) MarkInstallSucceeded(version string) {
	is.State = InstallStateInstalled
	is.Version = version
}

func (is *InstallStatus) MarkInstallRunning() {
	is.State = InstallStateInstalling
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Install is the Schema for the installs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=installs,scope=Namespaced
type Install struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstallSpec   `json:"spec,omitempty"`
	Status InstallStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InstallList contains a list of Install
type InstallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Install `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Install{}, &InstallList{})
}
