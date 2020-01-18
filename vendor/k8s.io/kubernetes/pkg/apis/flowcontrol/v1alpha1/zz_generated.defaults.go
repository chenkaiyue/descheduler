// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1alpha1.FlowSchema{}, func(obj interface{}) { SetObjectDefaults_FlowSchema(obj.(*v1alpha1.FlowSchema)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.FlowSchemaList{}, func(obj interface{}) { SetObjectDefaults_FlowSchemaList(obj.(*v1alpha1.FlowSchemaList)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.PriorityLevelConfiguration{}, func(obj interface{}) {
		SetObjectDefaults_PriorityLevelConfiguration(obj.(*v1alpha1.PriorityLevelConfiguration))
	})
	scheme.AddTypeDefaultingFunc(&v1alpha1.PriorityLevelConfigurationList{}, func(obj interface{}) {
		SetObjectDefaults_PriorityLevelConfigurationList(obj.(*v1alpha1.PriorityLevelConfigurationList))
	})
	return nil
}

func SetObjectDefaults_FlowSchema(in *v1alpha1.FlowSchema) {
	SetDefaults_FlowSchemaSpec(&in.Spec)
}

func SetObjectDefaults_FlowSchemaList(in *v1alpha1.FlowSchemaList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_FlowSchema(a)
	}
}

func SetObjectDefaults_PriorityLevelConfiguration(in *v1alpha1.PriorityLevelConfiguration) {
	if in.Spec.Limited != nil {
		SetDefaults_LimitedPriorityLevelConfiguration(in.Spec.Limited)
		if in.Spec.Limited.LimitResponse.Queuing != nil {
			SetDefaults_QueuingConfiguration(in.Spec.Limited.LimitResponse.Queuing)
		}
	}
}

func SetObjectDefaults_PriorityLevelConfigurationList(in *v1alpha1.PriorityLevelConfigurationList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PriorityLevelConfiguration(a)
	}
}