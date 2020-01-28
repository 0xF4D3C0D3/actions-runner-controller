/*
Copyright 2020 The actions-runner-controller authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RunnerSpec defines the desired state of Runner
type RunnerSpec struct {
	Repository string `json:"repository"`

	// +optional
	Image string `json:"image"`
}

// RunnerStatus defines the observed state of Runner
type RunnerStatus struct {
	Registration RunnerStatusRegistration `json:"registration"`
	Phase        string                   `json:"Phase"`
	Reason       string                   `json:"Reason"`
	Message      string                   `json:"Message"`
}

type RunnerStatusRegistration struct {
	Repository string      `json:"repository"`
	Token      string      `json:"token"`
	ExpiresAt  metav1.Time `json:"expiresAt"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Runner is the Schema for the runners API
type Runner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunnerSpec   `json:"spec,omitempty"`
	Status RunnerStatus `json:"status,omitempty"`
}

func (r Runner) IsRegisterable() bool {
	if r.Status.Registration.Repository != r.Spec.Repository {
		return false
	}

	if r.Status.Registration.Token == "" {
		return false
	}

	now := metav1.Now()
	if r.Status.Registration.ExpiresAt.Before(&now) {
		return false
	}

	return true
}

// +kubebuilder:object:root=true

// RunnerList contains a list of Runner
type RunnerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Runner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Runner{}, &RunnerList{})
}
