package render

import (
	operatorsv1alpha1 "github.com/vincent-pli/operator-sample-kustomize/pkg/apis/install/v1alpha1"
	"github.com/vincent-pli/operator-sample-kustomize/pkg/render/templates"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kustomize/v3/pkg/resource"
)

type renderFn func(*resource.Resource) (*unstructured.Unstructured, error)

type Renderer struct {
	cr *operatorsv1alpha1.Install
}

func NewRenderer(install *operatorsv1alpha1.Install) *Renderer {
	renderer := &Renderer{cr: install}
	return renderer
}

func (r *Renderer) Render() ([]*unstructured.Unstructured, error) {
	templates, err := templates.GetTemplateRenderer().GetTemplates(r.cr)
	if err != nil {
		return nil, err
	}

	uobjs := []*unstructured.Unstructured{}
	for _, template := range templates {
		uobjs = append(uobjs, &unstructured.Unstructured{Object: template.Map()})
	}
	return uobjs, nil
}
