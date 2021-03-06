// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

package policy

import (
	api "k8s.io/client-go/1.5/pkg/api"
	unversioned "k8s.io/client-go/1.5/pkg/api/unversioned"
	conversion "k8s.io/client-go/1.5/pkg/conversion"
	runtime "k8s.io/client-go/1.5/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_Eviction, InType: reflect.TypeOf(&Eviction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudget, InType: reflect.TypeOf(&PodDisruptionBudget{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudgetList, InType: reflect.TypeOf(&PodDisruptionBudgetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudgetSpec, InType: reflect.TypeOf(&PodDisruptionBudgetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudgetStatus, InType: reflect.TypeOf(&PodDisruptionBudgetStatus{})},
	)
}

func DeepCopy_policy_Eviction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Eviction)
		out := out.(*Eviction)
		out.TypeMeta = in.TypeMeta
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.DeleteOptions != nil {
			in, out := &in.DeleteOptions, &out.DeleteOptions
			*out = new(api.DeleteOptions)
			if err := api.DeepCopy_api_DeleteOptions(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.DeleteOptions = nil
		}
		return nil
	}
}

func DeepCopy_policy_PodDisruptionBudget(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudget)
		out := out.(*PodDisruptionBudget)
		out.TypeMeta = in.TypeMeta
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_policy_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_policy_PodDisruptionBudgetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudgetList)
		out := out.(*PodDisruptionBudgetList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PodDisruptionBudget, len(*in))
			for i := range *in {
				if err := DeepCopy_policy_PodDisruptionBudget(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_policy_PodDisruptionBudgetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudgetSpec)
		out := out.(*PodDisruptionBudgetSpec)
		out.MinAvailable = in.MinAvailable
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(unversioned.LabelSelector)
			if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		return nil
	}
}

func DeepCopy_policy_PodDisruptionBudgetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudgetStatus)
		out := out.(*PodDisruptionBudgetStatus)
		out.PodDisruptionAllowed = in.PodDisruptionAllowed
		out.CurrentHealthy = in.CurrentHealthy
		out.DesiredHealthy = in.DesiredHealthy
		out.ExpectedPods = in.ExpectedPods
		return nil
	}
}
