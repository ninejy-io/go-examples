package jobs

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	"k8s.io/client-go/util/retry"
)

type JobClient struct {
	client v1.JobInterface
	ns     string
}

func NewJobClient(clientset *kubernetes.Clientset, namespace string) *JobClient {
	return &JobClient{
		client: clientset.BatchV1().Jobs(namespace),
		ns:     namespace,
	}
}

func (j *JobClient) Create(name string) (*batchv1.Job, error) {
	job := &batchv1.Job{}

	return j.client.Create(context.TODO(), job, metav1.CreateOptions{})
}

func (j *JobClient) Update(name string) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		job, getErr := j.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		// job.Spec.Parallelism

		_, updateErr := j.client.Update(context.TODO(), job, metav1.UpdateOptions{})
		return updateErr
	})

	return retryErr
}

func (j *JobClient) Delete(name string) error {
	return j.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (j *JobClient) Get(name string) (*batchv1.Job, error) {
	return j.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (j *JobClient) List(labels ...string) ([]batchv1.Job, error) {
	var jobs *batchv1.JobList
	var err error

	if len(labels) == 0 {
		jobs, err = j.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		jobs, err = j.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return jobs.Items, nil
}
