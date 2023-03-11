// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2022. projectsveltos.io. All rights reserved.

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

var ClassifierReportFile = "../../config/crd/bases/lib.projectsveltos.io_classifierreports.yaml"
var ClassifierReportCRD = []byte(`---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: classifierreports.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: ClassifierReport
    listKind: ClassifierReportList
    plural: classifierreports
    singular: classifierreport
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ClassifierReport is the Schema for the classifierreports API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              classifierName:
                description: ClassifierName is the name of the Classifier instance
                  this report is for.
                type: string
              clusterName:
                description: ClusterName is the name of the Cluster this ClusterReport
                  is for.
                type: string
              clusterNamespace:
                description: ClusterNamespace is the namespace of the Cluster this
                  ClusterReport is for.
                type: string
              clusterType:
                description: ClusterType is the type of Cluster
                type: string
              match:
                description: Match indicates whether Cluster is currently a match
                  for the Classifier instance this report is for
                type: boolean
            required:
            - classifierName
            - clusterName
            - clusterNamespace
            - clusterType
            - match
            type: object
          status:
            description: ClassifierReportStatus defines the observed state of ClassifierReport
            properties:
              phase:
                description: Phase represents the current phase of report.
                enum:
                - WaitingForDelivery
                - Delivering
                - Processed
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`)
