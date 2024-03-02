package cronjob

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

type CronjobClient struct {
	client v1.CronJobInterface
	ns     string
}

func NewCronjobClient(clientset *kubernetes.Clientset, namespace string) *CronjobClient {
	return &CronjobClient{
		client: clientset.BatchV1().CronJobs(namespace),
		ns:     namespace,
	}
}

func (c *CronjobClient) Create(name string) {}

func (c *CronjobClient) Update(name string) {}

func (c *CronjobClient) Delete(name string) error {
	return c.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (c *CronjobClient) Get(name string) (*batchv1.CronJob, error) {
	return c.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *CronjobClient) List(labels ...string) ([]batchv1.CronJob, error) {
	var cronjobs *batchv1.CronJobList
	var err error

	if len(labels) == 0 {
		cronjobs, err = c.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		cronjobs, err = c.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return cronjobs.Items, nil
}
