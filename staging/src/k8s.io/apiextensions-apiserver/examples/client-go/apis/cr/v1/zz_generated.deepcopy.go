// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_v1_Example, InType: reflect.TypeOf(&Example{})},
		{Fn: DeepCopy_v1_ExampleList, InType: reflect.TypeOf(&ExampleList{})},
		{Fn: DeepCopy_v1_ExampleSpec, InType: reflect.TypeOf(&ExampleSpec{})},
		{Fn: DeepCopy_v1_ExampleStatus, InType: reflect.TypeOf(&ExampleStatus{})},
	}
}

// DeepCopy_v1_Example is an autogenerated deepcopy function.
func DeepCopy_v1_Example(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Example)
		out := out.(*Example)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		return nil
	}
}

// DeepCopy_v1_ExampleList is an autogenerated deepcopy function.
func DeepCopy_v1_ExampleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExampleList)
		out := out.(*ExampleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Example, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Example)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ExampleSpec is an autogenerated deepcopy function.
func DeepCopy_v1_ExampleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExampleSpec)
		out := out.(*ExampleSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1_ExampleStatus is an autogenerated deepcopy function.
func DeepCopy_v1_ExampleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExampleStatus)
		out := out.(*ExampleStatus)
		*out = *in
		return nil
	}
}