package install

import (
	"github.com/vincent-pli/operator-sample-kustomize/pkg/extension/imagereplacement"
)

func init() {
	activities = append(activities, imagereplacement.Configure)
}
