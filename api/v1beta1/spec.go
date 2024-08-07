/*
Copyright 2024. projectsveltos.io. All rights reserved.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
)

type Spec struct {
	// ClusterSelector identifies clusters to associate to
	// +optional
	ClusterSelector Selector `json:"clusterSelector,omitempty"`

	// ClusterRefs identifies clusters to associate to.
	// +optional
	ClusterRefs []corev1.ObjectReference `json:"clusterRefs,omitempty"`

	// MaxReplicas specifies the maximum number of clusters to be selected
	// from the pool matching the clusterSelector.
	MaxReplicas int `json:"maxReplicas,omitempty"`
}
