package v1

import (
	"istio.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:validation:Optional

//	限定请求来源：
//
// namespace 或者 特定工作负载
type From struct {
	//+kubebuilder:validation:Type=string
	//+kubebuilder:validation:Required
	Namespace string `json:"namespace,omitempty"`
	//+kubebuilder:validation:Optional
	WorkLoadSelector map[string]string `json:"workloadSelector,omitempty"`
}

type ByEgressGateway struct {

	//+kubebuilder:validation:Type=string
	//+kubebuilder:validation:Optional
	Name string `json:"name,omitempty"`
	//+kubebuilder:validation:Optional
	//+kubebuilder:validation:Maximum=65535
	Port uint32 `json:"port,omitempty"`
}

type HttpsUpgrade struct {
	Enabled *bool `json:"enabled,omitempty"`
	//+kubebuilder:validation:Maximum=65535
	Port uint32 `json:"port,omitempty"`
}

// Egress Service
type To struct {
	//+kubebuilder:validation:Type=string
	Name string `json:"name,omitempty"`
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:MinItems=1
	//+kubebuilder:validation:MaxItems=50
	Hosts []string `json:"hosts,omitempty"`
	//+kubebuilder:validation:Required
	Port *v1beta1.Port `json:"port,omitempty"`
	//+kubebuilder:validation:Required
	ByEgressGateway *ByEgressGateway `json:"byEgressGateway,omitempty"`
	HttpsUpgrade    *HttpsUpgrade    `json:"httpsUpgrade,omitempty"`
}

type EgressRule struct {
	//+kubebuilder:validation:Type=string
	//+kubebuilder:validation:optional
	Name string `json:"name,omitempty"`
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:MinItems=1
	From []*From `json:"from,omitempty"`
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:MinItems=1
	To []*To `json:"to,omitempty"`
}

// EgressTrafficPolicySpec defines the Spec of EgressTrafficPolicy
type EgressTrafficPolicySpec struct {
	//+kubebuilder:validation:Required
	ByEgressGateway *ByEgressGateway `json:"byEgressGateway,omitempty"`
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:MinItems=1
	EgressRules []*EgressRule `json:"egressRules,omitempty"`
}

// EgressTrafficPolicyStatus defines the observed state of EgressTrafficPolicy
type EgressTrafficPolicyStatus struct {
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="ByEgressGateway",type=string,JSONPath=`.spec.byEgressGateway.name`
//+kubebuilder:printcolumn:JSONPath=".status.message",name=status,type=string
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// ASMEgressTrafficPolicy is the Schema for the ASMEgressTrafficPolicy API
// +genclient
type ASMEgressTrafficPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	//+kubebuilder:validation:Required
	Spec   EgressTrafficPolicySpec   `json:"spec,omitempty"`
	Status EgressTrafficPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ASMEgressTrafficPolicyList contains a list of ASMEgressTrafficPolicy
type ASMEgressTrafficPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMEgressTrafficPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMEgressTrafficPolicy{}, &ASMEgressTrafficPolicyList{})
}
