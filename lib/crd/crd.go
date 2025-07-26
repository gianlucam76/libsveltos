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

//go:generate go run ../../generator.go

func GetClassifierCRDYAML() []byte {
	return ClassifierCRD
}

func GetClassifierReportCRDYAML() []byte {
	return ClassifierReportCRD
}

func GetDebuggingConfigurationCRDYAML() []byte {
	return DebuggingConfigurationCRD
}

func GetAccessRequestCRDYAML() []byte {
	return AccessRequestCRD
}

func GetSveltosClusterCRDYAML() []byte {
	return SveltosClusterCRD
}

func GetResourceSummaryCRDYAML() []byte {
	return ResourceSummaryCRD
}

func GetRoleRequestCRDYAML() []byte {
	return RoleRequestCRD
}

func GetClusterHealthCheckCRDYAML() []byte {
	return ClusterHealthCheckCRD
}

func GetHealthCheckCRDYAML() []byte {
	return HealthCheckCRD
}

func GetHealthCheckReportCRDYAML() []byte {
	return HealthCheckReportCRD
}

func GetEventSourceCRDYAML() []byte {
	return EventSourceCRD
}

func GetEventReportCRDYAML() []byte {
	return EventReportCRD
}

func GetReloaderCRDYAML() []byte {
	return ReloaderCRD
}

func GetReloaderReportCRDYAML() []byte {
	return ReloaderReportCRD
}

func GetClusterSetCRDYAML() []byte {
	return ClusterSetCRD
}

func GetSetCRDYAML() []byte {
	return SetCRD
}

func GetTechsupportCRDYAML() []byte {
	return TechsupportCRD
}

func GetConfigurationBundleCRDYAML() []byte {
	return ConfigurationBundleCRD
}

func GetConfigurationGroupCRDYAML() []byte {
	return ConfigurationGroupCRD
}

func GetSveltosLicenseCRDYAML() []byte {
	return SveltosLicenseCRD
}
