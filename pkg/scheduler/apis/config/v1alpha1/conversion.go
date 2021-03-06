/*
Copyright 2019 The Kubernetes Authors.

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
	"fmt"

	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/kube-scheduler/config/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
)

// Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in *v1alpha1.KubeSchedulerLeaderElectionConfiguration, out *config.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in, out, s); err != nil {
		return err
	}
	if len(in.ResourceNamespace) > 0 && len(in.LockObjectNamespace) == 0 {
		out.ResourceNamespace = in.ResourceNamespace
	} else if len(in.ResourceNamespace) == 0 && len(in.LockObjectNamespace) > 0 {
		out.ResourceNamespace = in.LockObjectNamespace
	} else if len(in.ResourceNamespace) > 0 && len(in.LockObjectNamespace) > 0 {
		if in.ResourceNamespace == in.LockObjectNamespace {
			out.ResourceNamespace = in.ResourceNamespace
		} else {
			return fmt.Errorf("ResourceNamespace and LockObjectNamespace are both set and do not match, ResourceNamespace: %s, LockObjectNamespace: %s", in.ResourceNamespace, in.LockObjectNamespace)
		}
	}

	if len(in.ResourceName) > 0 && len(in.LockObjectName) == 0 {
		out.ResourceName = in.ResourceName
	} else if len(in.ResourceName) == 0 && len(in.LockObjectName) > 0 {
		out.ResourceName = in.LockObjectName
	} else if len(in.ResourceName) > 0 && len(in.LockObjectName) > 0 {
		if in.ResourceName == in.LockObjectName {
			out.ResourceName = in.ResourceName
		} else {
			return fmt.Errorf("ResourceName and LockObjectName are both set and do not match, ResourceName: %s, LockObjectName: %s", in.ResourceName, in.LockObjectName)
		}
	}
	return nil
}

// Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *config.KubeSchedulerLeaderElectionConfiguration, out *v1alpha1.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := autoConvert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in, out, s); err != nil {
		return err
	}
	out.ResourceNamespace = in.ResourceNamespace
	out.LockObjectNamespace = in.ResourceNamespace
	out.ResourceName = in.ResourceName
	out.LockObjectName = in.ResourceName
	return nil
}
