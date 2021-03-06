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
package imagereplacement

import (
	operatorv1alpha1 "github.com/vincent-pli/operator-sample-kustomize/pkg/apis/install/v1alpha1"
	"github.com/vincent-pli/operator-sample-kustomize/pkg/extension/common"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	extension = common.Extension{
		Transformers: []common.Transformer{egress},
	}
	log      = logf.Log.WithName("image-replacement")
	scheme   *runtime.Scheme
	instance *operatorv1alpha1.Install
)

func Configure(c client.Client, s *runtime.Scheme, install *operatorv1alpha1.Install) (*common.Extension, error) {
	if install.Spec.Registry != nil {
		log.Info("Enalbe extension: Imagereplacement")
		scheme = s
		instance = install
		return &extension, nil
	}

	return nil, nil
}

func egress(u *unstructured.Unstructured) error {
	if u.GetKind() == "Deployment" {
		var deploy = &appsv1.Deployment{}
		if err := scheme.Convert(u, deploy, nil); err != nil {
			return err
		}
		registry := instance.Spec.Registry
		err := UpdateDeployment(deploy, registry, log)
		if err != nil {
			return err
		}
		if err := scheme.Convert(deploy, u, nil); err != nil {
			return err
		}
	}
	return nil
}
