package deployer

import (
	"context"
	"reflect"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func Deploy(c runtimeclient.Client, obj *unstructured.Unstructured) error {
	found := &unstructured.Unstructured{}
	found.SetGroupVersionKind(obj.GroupVersionKind())
	err := c.Get(context.TODO(), types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}, found)
	if err != nil {
		if errors.IsNotFound(err) {
			return c.Create(context.TODO(), obj)
		}
		return err
	}

	if found.GetKind() != "Deployment" {
		return nil
	}

	oldSpec, oldSpecFound := found.Object["spec"]
	newSpec, newSpecFound := obj.Object["spec"]
	if !oldSpecFound || !newSpecFound {
		return nil
	}
	if !reflect.DeepEqual(oldSpec, newSpec) {
		newObj := found.DeepCopy()
		newObj.Object["spec"] = newSpec
		return c.Update(context.TODO(), newObj)
	}
	return nil
}
