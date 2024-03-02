package daemonset

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type DaemonSetClient struct {
	client v1.DaemonSetInterface
	ns     string
}

func NewDaemonsetClient(clientset *kubernetes.Clientset, namespace string) *DaemonSetClient {
	return &DaemonSetClient{
		client: clientset.AppsV1().DaemonSets(namespace),
		ns:     namespace,
	}
}

func (d *DaemonSetClient) Create(name string) {}

func (d *DaemonSetClient) Update(name string) {}

func (d *DaemonSetClient) Delete(name string) error {
	return d.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (d *DaemonSetClient) Get(name string) (*appsv1.DaemonSet, error) {
	return d.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (d *DaemonSetClient) List(labels ...string) ([]appsv1.DaemonSet, error) {
	var ds *appsv1.DaemonSetList
	var err error

	if len(labels) == 0 {
		ds, err = d.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		ds, err = d.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return ds.Items, nil
}
