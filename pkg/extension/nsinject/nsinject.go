package nsinject

import (
	"os"
	"strings"

	operatorsv1alpha1 "github.com/vincent-pli/operator-sample-kustomize/pkg/apis/install/v1alpha1"
	"github.com/vincent-pli/operator-sample-kustomize/pkg/extension/common"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	extension = common.Extension{
		Transformers: []common.Transformer{egress},
	}

	namespace string
)

func Configure(c client.Client, s *runtime.Scheme, install *operatorsv1alpha1.Install) (*common.Extension, error) {
	if &install.Spec.TargetNamespace != nil {
		namespace = install.Spec.TargetNamespace
		return &extension, nil
	}

	return nil, nil
}

func egress(u *unstructured.Unstructured) error {
	switch strings.ToLower(u.GetKind()) {
	case "namespace":
		u.SetName(namespace)
	case "clusterrolebinding":
		subjects, _, _ := unstructured.NestedFieldNoCopy(u.Object, "subjects")
		for _, subject := range subjects.([]interface{}) {
			m := subject.(map[string]interface{})
			if _, ok := m["namespace"]; ok {
				m["namespace"] = namespace
			}
		}
	}
	if !isClusterScoped(u.GetKind()) {
		u.SetNamespace(namespace)
	}

	return nil
}

func resolveEnv(x string) string {
	if len(x) > 1 && x[:1] == "$" {
		return os.Getenv(x[1:])
	}
	return x
}

func isClusterScoped(kind string) bool {
	// TODO: something more clever using !APIResource.Namespaced maybe?
	switch strings.ToLower(kind) {
	case "componentstatus",
		"namespace",
		"node",
		"persistentvolume",
		"mutatingwebhookconfiguration",
		"validatingwebhookconfiguration",
		"customresourcedefinition",
		"apiservice",
		"meshpolicy",
		"tokenreview",
		"selfsubjectaccessreview",
		"selfsubjectrulesreview",
		"subjectaccessreview",
		"certificatesigningrequest",
		"podsecuritypolicy",
		"clusterrolebinding",
		"clusterrole",
		"priorityclass",
		"storageclass",
		"volumeattachment":
		return true
	}
	return false
}
