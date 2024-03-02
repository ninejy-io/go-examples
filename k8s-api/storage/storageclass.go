package storage

import (
	"context"

	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

type StorageClassClient struct {
	client v1.StorageClassInterface
}

func NewStorageClassClient(clientset *kubernetes.Clientset, namespace string) *StorageClassClient {
	return &StorageClassClient{
		client: clientset.StorageV1().StorageClasses(),
	}
}

func (s *StorageClassClient) Create(name string) {}

func (s *StorageClassClient) Update(name string) {}

func (s *StorageClassClient) Delete(name string) error {
	return s.client.Delete(context.Background(), name, metav1.DeleteOptions{})
}

func (s *StorageClassClient) Get(name string) (*storagev1.StorageClass, error) {
	return s.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (s *StorageClassClient) List(labels ...string) ([]storagev1.StorageClass, error) {
	var storages *storagev1.StorageClassList
	var err error

	if len(labels) == 0 {
		storages, err = s.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		storages, err = s.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if len(labels) == 0 {
		return nil, err
	}
	return storages.Items, nil
}
