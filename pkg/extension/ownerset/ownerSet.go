package ownerset

import (
	operatorsv1alpha1 "github.com/vincent-pli/operator-sample-kustomize/pkg/apis/install/v1alpha1"
	"github.com/vincent-pli/operator-sample-kustomize/pkg/extension/common"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	extension = common.Extension{
		Transformers: []common.Transformer{egress},
	}
	log      = logf.Log.WithName("owner-setting")
	scheme   *runtime.Scheme
	instance *operatorsv1alpha1.Install
)

func Configure(c client.Client, s *runtime.Scheme, install *operatorsv1alpha1.Install) (*common.Extension, error) {
	if install.Spec.SetOwner {
		log.Info("Enable extension: ownerset")
		scheme = s
		instance = install
		return &extension, nil
	}

	return nil, nil
}

func egress(u *unstructured.Unstructured) error {
	if err := controllerutil.SetControllerReference(instance, u, scheme); err != nil {
		return err
	}

	return nil
}
