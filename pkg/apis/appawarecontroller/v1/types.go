package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ForecastPolicy is a specification for a ForecastPolicy resource
type ForecastPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ForecastPolicySpec `json:"spec"`
}

// ForecastPolicySpec is the spec for a ForecastPolicy resource
type ForecastPolicySpec struct {
	ForecastWindow *int32             `json:"forecastWindow"`
	ActionType     string             `json:"actionType"`
	Selector       []ForecastSelector `json:"selector"`
}

// ForecastSelector is the selector for a ForecastPolicy resource
type ForecastSelector struct {
	MatchLabels []MatchLabel `json:"matchLabels"`
}

// MatchLabel is the matchLabels for a ForecastSelector resource
type MatchLabel struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ForecastPolicyList is a list of ForecastPolicy resources
type ForecastPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ForecastPolicy `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceRecommendation is a specification for a ResourceRecommendation resource
type ResourceRecommendation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ResourceRecommendationSpec `json:"spec"`
}

type ResourceRecommendationSpec struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Resources []ResourcePlacement `json:"resources"`
}

type ResourcePlacement struct {
	Name string `json:"name"`
	Placements []Placement `json:"placements"`
}

type Placement struct {
	FlavorName string `json:"flavorName"`
	Amount *int32 `json:"amount"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceRecommendationList is a list of ResourceRecommendation resources
type ResourceRecommendationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ResourceRecommendation `json:"items"`
}
