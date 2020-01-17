package install

import (
	"github.com/vincent-pli/operator-sample-kustomize/pkg/extension/ownerset"
)

func init() {
	activities = append(activities, ownerset.Configure)
}
