package util

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	"testing"
)

func TestRemoveDuplicatedVolumes(t *testing.T) {
	volumes := []v1.Volume{
		v1.Volume{
			Name:         "volume1",
			VolumeSource: v1.VolumeSource{},
		},
		v1.Volume{
			Name:         "volume2",
			VolumeSource: v1.VolumeSource{},
		},
		v1.Volume{
			Name:         "volume1",
			VolumeSource: v1.VolumeSource{},
		},
	}

	assert.Len(t, RemoveDuplicatedVolumes(volumes), 2)
	assert.Equal(t, "volume1", volumes[0].Name)
	assert.Equal(t, "volume2", volumes[1].Name)
}

func TestRemoveDuplicatedVolumeMounts(t *testing.T) {
	volumeMounts := []v1.VolumeMount{
		v1.VolumeMount{
			Name:     "data1",
			ReadOnly: false,
		},
		v1.VolumeMount{
			Name:     "data2",
			ReadOnly: false,
		},
		v1.VolumeMount{
			Name:     "data1",
			ReadOnly: false,
		},
	}

	assert.Len(t, RemoveDuplicatedVolumeMounts(volumeMounts), 2)
	assert.Equal(t, "data1", volumeMounts[0].Name)
	assert.Equal(t, "data2", volumeMounts[1].Name)
}
