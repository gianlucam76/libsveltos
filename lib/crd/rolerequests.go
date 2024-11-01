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

var RoleRequestFile = "../../manifests/apiextensions.k8s.io_v1_customresourcedefinition_rolerequests.lib.projectsveltos.io.yaml"
var RoleRequestCRD = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    cert-manager.io/inject-ca-from: projectsveltos/projectsveltos-serving-cert
    controller-gen.kubebuilder.io/version: v0.16.5
  name: rolerequests.lib.projectsveltos.io
spec:
  conversion:
    strategy: Webhook
    webhook:
      clientConfig:
        service:
          name: webhook-service
          namespace: projectsveltos
          path: /convert
      conversionReviewVersions:
      - v1
  group: lib.projectsveltos.io
  names:
    kind: RoleRequest
    listKind: RoleRequestList
    plural: rolerequests
    singular: rolerequest
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: RoleRequest is the Schema for the rolerequest API
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
            description: RoleRequestSpec defines the desired state of RoleRequest
            properties:
              clusterSelector:
                description: |-
                  ClusterSelector identifies clusters where permissions requestes
                  in this instance will be granted (Deprecated use selector instead)
                type: string
              expirationSeconds:
                description: |-
                  ExpirationSeconds is the requested duration of validity of the TokenRequest
                  associated to ServiceAccount. If not specified, default value is used
                format: int64
                type: integer
              roleRefs:
                description: |-
                  RoleRefs references all the Secret/ConfigMaps containing kubernetes
                  Roles/ClusterRoles that need to be deployed in the matching clusters.
                items:
                  description: |-
                    PolicyRef specifies a resource containing one or more policy
                    to deploy in matching Clusters.
                  properties:
                    kind:
                      description: 'Kind of the resource. Supported kinds are: Secrets
                        and ConfigMaps.'
                      enum:
                      - Secret
                      - ConfigMap
                      type: string
                    name:
                      description: Name of the referenced resource.
                      minLength: 1
                      type: string
                    namespace:
                      description: |-
                        Namespace of the referenced resource.
                        Namespace can be left empty. In such a case, namespace will
                        be implicit set to cluster's namespace.
                      type: string
                  required:
                  - kind
                  - name
                  - namespace
                  type: object
                type: array
              serviceAccountName:
                description: |-
                  ServiceAccountName is the name of the ServiceAccount representing a tenant admin for which
                  those permissions are requested
                type: string
              serviceAccountNamespace:
                description: |-
                  ServiceAccountNamespace is the name of the ServiceAccount representing a tenant admin
                  for which those permissions are requested
                type: string
            required:
            - clusterSelector
            - serviceAccountName
            - serviceAccountNamespace
            type: object
          status:
            description: RoleRequestStatus defines the status of RoleRequest
            properties:
              clusterInfo:
                description: |-
                  ClusterInfo represents the hash of the ClusterRoles/Roles deployed in
                  a matching cluster for the admin.
                items:
                  properties:
                    cluster:
                      description: Cluster references the Cluster
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
                    failureMessage:
                      description: FailureMessage provides more information about
                        the error.
                      type: string
                    hash:
                      description: |-
                        Hash represents the hash of the Classifier currently deployed
                        in the Cluster
                      format: byte
                      type: string
                    status:
                      description: Status represents the state of the feature in the
                        workload cluster
                      enum:
                      - Provisioning
                      - Provisioned
                      - Failed
                      - Removing
                      - Removed
                      type: string
                  required:
                  - cluster
                  - hash
                  type: object
                type: array
              failureMessage:
                description: FailureMessage provides more information if an error
                  occurs.
                type: string
              matchingClusters:
                description: |-
                  MatchingClusterRefs reference all the cluster currently matching
                  RoleRequest ClusterSelector
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
    storage: false
    subresources:
      status: {}
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: RoleRequest is the Schema for the rolerequest API
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
            description: RoleRequestSpec defines the desired state of RoleRequest
            properties:
              clusterSelector:
                description: Selector identifies clusters to associate to.
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
              expirationSeconds:
                description: |-
                  ExpirationSeconds is the requested duration of validity of the TokenRequest
                  associated to ServiceAccount. If not specified, default value is used
                format: int64
                type: integer
              roleRefs:
                description: |-
                  RoleRefs references all the Secret/ConfigMaps containing kubernetes
                  Roles/ClusterRoles that need to be deployed in the matching clusters.
                items:
                  description: |-
                    PolicyRef specifies a resource containing one or more policy
                    to deploy in matching Clusters.
                  properties:
                    kind:
                      description: 'Kind of the resource. Supported kinds are: Secrets
                        and ConfigMaps.'
                      enum:
                      - Secret
                      - ConfigMap
                      type: string
                    name:
                      description: Name of the referenced resource.
                      minLength: 1
                      type: string
                    namespace:
                      description: |-
                        Namespace of the referenced resource.
                        Namespace can be left empty. In such a case, namespace will
                        be implicit set to cluster's namespace.
                      type: string
                  required:
                  - kind
                  - name
                  - namespace
                  type: object
                type: array
              serviceAccountName:
                description: |-
                  ServiceAccountName is the name of the ServiceAccount representing a tenant admin for which
                  those permissions are requested
                type: string
              serviceAccountNamespace:
                description: |-
                  ServiceAccountNamespace is the name of the ServiceAccount representing a tenant admin
                  for which those permissions are requested
                type: string
            required:
            - serviceAccountName
            - serviceAccountNamespace
            type: object
          status:
            description: RoleRequestStatus defines the status of RoleRequest
            properties:
              clusterInfo:
                description: |-
                  ClusterInfo represents the hash of the ClusterRoles/Roles deployed in
                  a matching cluster for the admin.
                items:
                  properties:
                    cluster:
                      description: Cluster references the Cluster
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
                    failureMessage:
                      description: FailureMessage provides more information about
                        the error.
                      type: string
                    hash:
                      description: |-
                        Hash represents the hash of the Classifier currently deployed
                        in the Cluster
                      format: byte
                      type: string
                    status:
                      description: Status represents the state of the feature in the
                        workload cluster
                      enum:
                      - Provisioning
                      - Provisioned
                      - Failed
                      - Removing
                      - Removed
                      type: string
                  required:
                  - cluster
                  - hash
                  type: object
                type: array
              failureMessage:
                description: FailureMessage provides more information if an error
                  occurs.
                type: string
              matchingClusters:
                description: |-
                  MatchingClusterRefs reference all the cluster currently matching
                  RoleRequest ClusterSelector
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
