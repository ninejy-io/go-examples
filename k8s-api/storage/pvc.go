package storage

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/util/retry"
)

type PersistentVolumeClaimClient struct {
	client v1.PersistentVolumeClaimInterface
	ns     string
}

func NewPersistentVolumeClaimClient(clientset *kubernetes.Clientset, namespace string) *PersistentVolumeClaimClient {
	return &PersistentVolumeClaimClient{
		client: clientset.CoreV1().PersistentVolumeClaims(namespace),
		ns:     namespace,
	}
}

func (p *PersistentVolumeClaimClient) Create(name string) (*corev1.PersistentVolumeClaim, error) {
	pvc := &corev1.PersistentVolumeClaim{}

	return p.client.Create(context.TODO(), pvc, metav1.CreateOptions{})
}

func (p *PersistentVolumeClaimClient) Update(name string) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		pvc, getErr := p.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		// pvc.Spec.Resources

		_, updateErr := p.client.Update(context.TODO(), pvc, metav1.UpdateOptions{})
		return updateErr
	})

	return retryErr
}

func (p *PersistentVolumeClaimClient) Delete(name string) error {
	return p.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (p *PersistentVolumeClaimClient) Get(name string) (*corev1.PersistentVolumeClaim, error) {
	return p.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (p *PersistentVolumeClaimClient) List(labels ...string) ([]corev1.PersistentVolumeClaim, error) {
	var pvcs *corev1.PersistentVolumeClaimList
	var err error

	if len(labels) == 0 {
		pvcs, err = p.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		pvcs, err = p.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return pvcs.Items, nil
}
