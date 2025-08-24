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

package logsettings_test

import (
	"flag"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	libsveltosv1beta1 "github.com/projectsveltos/libsveltos/api/v1beta1"
	logs "github.com/projectsveltos/libsveltos/lib/logsettings"
)

var _ = Describe("Logsettings", func() {
	It("Should change log level appropriately", func() {
		conf := &libsveltosv1beta1.DebuggingConfiguration{
			ObjectMeta: metav1.ObjectMeta{
				Name: "default",
			},
			Spec: libsveltosv1beta1.DebuggingConfigurationSpec{
				Configuration: []libsveltosv1beta1.ComponentConfiguration{
					{Component: libsveltosv1beta1.ComponentAddonManager, LogLevel: libsveltosv1beta1.LogLevelDebug},
				},
			},
		}

		logs.UpdateLogLevel(conf)
		f := flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(logs.LogDebug)))

		conf.Spec.Configuration = []libsveltosv1beta1.ComponentConfiguration{
			{Component: libsveltosv1beta1.ComponentAddonManager, LogLevel: libsveltosv1beta1.LogLevelInfo},
		}

		logs.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(logs.LogInfo)))

		conf.Spec.Configuration = []libsveltosv1beta1.ComponentConfiguration{
			{Component: libsveltosv1beta1.ComponentAddonManager, LogLevel: libsveltosv1beta1.LogLevelVerbose},
		}

		logs.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(logs.LogVerbose)))

		newDebugValue := 8
		instance.SetDebugValue(newDebugValue)
		conf.Spec.Configuration = []libsveltosv1beta1.ComponentConfiguration{
			{Component: libsveltosv1beta1.ComponentAddonManager, LogLevel: libsveltosv1beta1.LogLevelDebug},
		}

		logs.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(newDebugValue)))

		newInfoValue := 5
		instance.SetInfoValue(newInfoValue)
		conf.Spec.Configuration = []libsveltosv1beta1.ComponentConfiguration{
			{Component: libsveltosv1beta1.ComponentAddonManager, LogLevel: libsveltosv1beta1.LogLevelInfo},
		}

		logs.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(newInfoValue)))
	})
})
