// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package vsphereproviderconfig

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereClusterProviderConfig) DeepCopyInto(out *VsphereClusterProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereClusterProviderConfig.
func (in *VsphereClusterProviderConfig) DeepCopy() *VsphereClusterProviderConfig {
	if in == nil {
		return nil
	}
	out := new(VsphereClusterProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereClusterProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereMachineProviderConfig) DeepCopyInto(out *VsphereMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.TerraformVariables != nil {
		in, out := &in.TerraformVariables, &out.TerraformVariables
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereMachineProviderConfig.
func (in *VsphereMachineProviderConfig) DeepCopy() *VsphereMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(VsphereMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereMachineProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}