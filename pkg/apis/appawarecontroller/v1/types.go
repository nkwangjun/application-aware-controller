package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppawareHorizontalPodAutoscaler is a specification for a AppawareHorizontalPodAutoscaler resource
type AppawareHorizontalPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppawareHorizontalPodAutoscalerSpec   `json:"spec"`
	Status AppawareHorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

type AppawareHorizontalPodAutoscalerSpec struct {
	ScaleTargetRef ScaleTargetRef `json:"scaleTargetRef"`
	ScaleMode      string         `json:"scaleMode"`
	ForecastWindow *int32         `json:"forecastWindow"`
	Jobs           []Job          `json:"jobs,omitempty"`
}

type ScaleTargetRef struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
}

type Job struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	// job will only run once if enabled.
	RunOnce    bool  `json:"runOnce,omitempty"`
	TargetSize int32 `json:"targetSize"`
}

// AppawareHorizontalPodAutoscalerStatus defines the observed state of AppawareHorizontalPodAutoscaler
type AppawareHorizontalPodAutoscalerStatus struct {
	JobStatus []JobStatus `json:"jobstatus,omitempty"`
}

type JobState string

const (
	Succeed   JobState = "Succeed"
	Failed    JobState = "Failed"
	Submitted JobState = "Submitted"
)

type JobStatus struct {
	JobId         string      `json:"jobId"`
	State         JobState    `json:"state"`
	LastProbeTime metav1.Time `json:"lastProbeTime"`
	// +optional
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppawareHorizontalPodAutoscalerList is a list of AppawareHorizontalPodAutoscaler resources
type AppawareHorizontalPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AppawareHorizontalPodAutoscaler `json:"items"`
}
