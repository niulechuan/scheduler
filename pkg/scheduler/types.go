package scheduler

import (
	v1 "k8s.io/api/core/v1"
)

type VolumeScheduler interface {
	Filter(existingLocalVolume []string, unboundPVCs []*v1.PersistentVolumeClaim, node *v1.Node) (bool, error)
	CSIDriverName() string
}
