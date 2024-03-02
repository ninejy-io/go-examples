package storage

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PersistentVolumeClient struct {
	client v1.PersistentVolumeInterface
}

func NewPersistentVolumeClient(clientset *kubernetes.Clientset) *PersistentVolumeClient {
	return &PersistentVolumeClient{
		client: clientset.CoreV1().PersistentVolumes(),
	}
}

func (p *PersistentVolumeClient) Create(name string) {}

func (p *PersistentVolumeClient) Update(name string) {}

func (p *PersistentVolumeClient) Delete(name string) error {
	return p.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (p *PersistentVolumeClient) Get(name string) (*corev1.PersistentVolume, error) {
	return p.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (p *PersistentVolumeClient) List(labels ...string) ([]corev1.PersistentVolume, error) {
	var pvs *corev1.PersistentVolumeList
	var err error

	if len(labels) == 0 {
		pvs, err = p.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		pvs, err = p.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return pvs.Items, nil
}
