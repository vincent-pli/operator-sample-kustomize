/*
Copyright 2019 The Knative Authors

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
package common

import (
	operatorsv1alpha1 "github.com/vincent-pli/operator-sample-kustomize/pkg/apis/install/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("common")

type Activities []func(client.Client, *runtime.Scheme, *operatorsv1alpha1.Install) (*Extension, error)
type Extender func(*operatorsv1alpha1.Install) error
type Transformer func(u *unstructured.Unstructured) error
type Extensions []Extension
type Extension struct {
	Transformers []Transformer
	PreInstalls  []Extender
	PostInstalls []Extender
}

func (activities Activities) Extend(c client.Client, scheme *runtime.Scheme, install *operatorsv1alpha1.Install) (Extensions, error) {
	result := Extensions{}
	for _, fn := range activities {
		ext, err := fn(c, scheme, install)
		if err != nil {
			return result, err
		}
		if ext != nil {
			result = append(result, *ext)
		}
	}
	return result, nil
}

func (exts Extensions) generate(install *operatorsv1alpha1.Install) []Transformer {
	result := []Transformer{}
	for _, extension := range exts {
		result = append(result, extension.Transformers...)
	}
	// Transformer will run in order, so can add some more Transformer here
	return append(result, func(u *unstructured.Unstructured) error {
		return nil
	})
}

func (exts Extensions) Transformer(resources []*unstructured.Unstructured, install *operatorsv1alpha1.Install) ([]*unstructured.Unstructured, error) {
	transformers := exts.generate(install)
	var results []*unstructured.Unstructured
	for i := 0; i < len(resources); i++ {
		spec := resources[i].DeepCopy()
		for _, transform := range transformers {
			err := transform(spec)
			if err != nil {
				return nil, err
			}
		}
		results = append(results, spec)
	}
	return results, nil
}

func (exts Extensions) PreInstall(install *operatorsv1alpha1.Install) error {
	for _, extension := range exts {
		for _, f := range extension.PreInstalls {
			if err := f(install); err != nil {
				return err
			}
		}
	}
	return nil
}

func (exts Extensions) PostInstall(install *operatorsv1alpha1.Install) error {
	for _, extension := range exts {
		for _, f := range extension.PostInstalls {
			if err := f(install); err != nil {
				return err
			}
		}
	}
	return nil
}
