package install

import (
	"github.com/vincent-pli/operator-sample-kustomize/pkg/extension/nsinject"
)

func init() {
	activities = append(activities, nsinject.Configure)
}
