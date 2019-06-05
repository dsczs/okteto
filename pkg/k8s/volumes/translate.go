package volumes

import (
	"fmt"

	"github.com/okteto/app/cli/pkg/model"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	oktetoVolumeTemplate     = "okteto-%s"
	oktetoVolumeDataTemplate = "okteto-%s-%d"
)

func translate(name string) *apiv1.PersistentVolumeClaim {
	quantDisk, _ := resource.ParseQuantity("10Gi")
	return &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{apiv1.ReadWriteOnce},
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					"storage": quantDisk,
				},
			},
		},
	}
}

//GetVolumeName returns the okteto volume name for a given dev environment
func GetVolumeName(dev *model.Dev) string {
	return fmt.Sprintf(oktetoVolumeTemplate, dev.Name)
}

//GetVolumeDataName returns the okteto volume name for a given dev environment
func GetVolumeDataName(dev *model.Dev, i int) string {
	return fmt.Sprintf(oktetoVolumeDataTemplate, dev.Name, i)
}