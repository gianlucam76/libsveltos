// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2022-23. projectsveltos.io. All rights reserved.

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
package crd

var ClusterSetFile = "../../manifests/apiextensions.k8s.io_v1_customresourcedefinition_clustersets.lib.projectsveltos.io.yaml"
var ClusterSetCRD = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.17.3
  name: clustersets.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: ClusterSet
    listKind: ClusterSetList
    plural: clustersets
    singular: clusterset
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ClusterSet is the Schema for the clustersets API
        properties:
          apiVersion:
            description: |-
              APIVersion defines the versioned schema of this representation of an object.
              Servers should convert recognized schemas to the latest internal value, and
              may reject unrecognized values.
              More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
            type: string
          kind:
            description: |-
              Kind is a string value representing the REST resource this object represents.
              Servers may infer this from the endpoint the client submits requests to.
              Cannot be updated.
              In CamelCase.
              More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
            type: string
          metadata:
            type: object
          spec:
            properties:
              clusterRefs:
                description: ClusterRefs identifies clusters to associate to.
                items:
                  description: ObjectReference contains enough information to let
                    you inspect or modify the referred object.
                  properties:
                    apiVersion:
                      description: API version of the referent.
                      type: string
                    fieldPath:
                      description: |-
                        If referring to a piece of an object instead of an entire object, this string
                        should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
                        For example, if the object reference is to a container within a pod, this would take on a value like:
                        "spec.containers{name}" (where "name" refers to the name of the container that triggered
                        the event) or if no container name is specified "spec.containers[2]" (container with
                        index 2 in this pod). This syntax is chosen only to have some well-defined way of
                        referencing a part of an object.
                      type: string
                    kind:
                      description: |-
                        Kind of the referent.
                        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
                      type: string
                    name:
                      description: |-
                        Name of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                      type: string
                    namespace:
                      description: |-
                        Namespace of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
                      type: string
                    resourceVersion:
                      description: |-
                        Specific resourceVersion to which this reference is made, if any.
                        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
                      type: string
                    uid:
                      description: |-
                        UID of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
                      type: string
                  type: object
                  x-kubernetes-map-type: atomic
                type: array
              clusterSelector:
                description: ClusterSelector identifies clusters to associate to
                properties:
                  matchExpressions:
                    description: matchExpressions is a list of label selector requirements.
                      The requirements are ANDed.
                    items:
                      description: |-
                        A label selector requirement is a selector that contains values, a key, and an operator that
                        relates the key and values.
                      properties:
                        key:
                          description: key is the label key that the selector applies
                            to.
                          type: string
                        operator:
                          description: |-
                            operator represents a key's relationship to a set of values.
                            Valid operators are In, NotIn, Exists and DoesNotExist.
                          type: string
                        values:
                          description: |-
                            values is an array of string values. If the operator is In or NotIn,
                            the values array must be non-empty. If the operator is Exists or DoesNotExist,
                            the values array must be empty. This array is replaced during a strategic
                            merge patch.
                          items:
                            type: string
                          type: array
                          x-kubernetes-list-type: atomic
                      required:
                      - key
                      - operator
                      type: object
                    type: array
                    x-kubernetes-list-type: atomic
                  matchLabels:
                    additionalProperties:
                      type: string
                    description: |-
                      matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
                      map is equivalent to an element of matchExpressions, whose key field is "key", the
                      operator is "In", and the values array contains only "value". The requirements are ANDed.
                    type: object
                type: object
                x-kubernetes-map-type: atomic
              maxReplicas:
                description: |-
                  MaxReplicas specifies the maximum number of clusters to be selected
                  from the pool matching the clusterSelector.
                type: integer
            type: object
          status:
            description: Status defines the observed state of ClusterSet/Set
            properties:
              matchingClusterRefs:
                description: |-
                  MatchingClusterRefs reference all the clusters currently matching
                  ClusterSet/Set ClusterSelector
                items:
                  description: ObjectReference contains enough information to let
                    you inspect or modify the referred object.
                  properties:
                    apiVersion:
                      description: API version of the referent.
                      type: string
                    fieldPath:
                      description: |-
                        If referring to a piece of an object instead of an entire object, this string
                        should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
                        For example, if the object reference is to a container within a pod, this would take on a value like:
                        "spec.containers{name}" (where "name" refers to the name of the container that triggered
                        the event) or if no container name is specified "spec.containers[2]" (container with
                        index 2 in this pod). This syntax is chosen only to have some well-defined way of
                        referencing a part of an object.
                      type: string
                    kind:
                      description: |-
                        Kind of the referent.
                        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
                      type: string
                    name:
                      description: |-
                        Name of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                      type: string
                    namespace:
                      description: |-
                        Namespace of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
                      type: string
                    resourceVersion:
                      description: |-
                        Specific resourceVersion to which this reference is made, if any.
                        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
                      type: string
                    uid:
                      description: |-
                        UID of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
                      type: string
                  type: object
                  x-kubernetes-map-type: atomic
                type: array
              selectedClusterRefs:
                description: |-
                  SelectedClusters reference all the cluster currently selected among
                  all the ones matching
                items:
                  description: ObjectReference contains enough information to let
                    you inspect or modify the referred object.
                  properties:
                    apiVersion:
                      description: API version of the referent.
                      type: string
                    fieldPath:
                      description: |-
                        If referring to a piece of an object instead of an entire object, this string
                        should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
                        For example, if the object reference is to a container within a pod, this would take on a value like:
                        "spec.containers{name}" (where "name" refers to the name of the container that triggered
                        the event) or if no container name is specified "spec.containers[2]" (container with
                        index 2 in this pod). This syntax is chosen only to have some well-defined way of
                        referencing a part of an object.
                      type: string
                    kind:
                      description: |-
                        Kind of the referent.
                        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
                      type: string
                    name:
                      description: |-
                        Name of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                      type: string
                    namespace:
                      description: |-
                        Namespace of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
                      type: string
                    resourceVersion:
                      description: |-
                        Specific resourceVersion to which this reference is made, if any.
                        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
                      type: string
                    uid:
                      description: |-
                        UID of the referent.
                        More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
                      type: string
                  type: object
                  x-kubernetes-map-type: atomic
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`)
