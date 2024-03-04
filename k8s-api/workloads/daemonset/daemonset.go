package daemonset

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"
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

func (d *DaemonSetClient) Create(name string) (*appsv1.DaemonSet, error) {
	ds := &appsv1.DaemonSet{}

	return d.client.Create(context.TODO(), ds, metav1.CreateOptions{})
}

func (d *DaemonSetClient) Update(name string) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		ds, getErr := d.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		// ds.Spec.Selector

		_, updateErr := d.client.Update(context.TODO(), ds, metav1.UpdateOptions{})
		return updateErr
	})

	return retryErr
}

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
