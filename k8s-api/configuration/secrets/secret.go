package secrets

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type SecretClient struct {
	client v1.SecretInterface
	ns     string
}

func NewSecretClient(clientset *kubernetes.Clientset, namespace string) *SecretClient {
	return &SecretClient{
		client: clientset.CoreV1().Secrets(namespace),
		ns:     namespace,
	}
}

func (s *SecretClient) Create(name string) {}

func (s *SecretClient) Update(name string) {}

func (s *SecretClient) Delete(name string) error {
	return s.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (s *SecretClient) Get(name string) (*corev1.Secret, error) {
	return s.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (s *SecretClient) List(labels ...string) ([]corev1.Secret, error) {
	var secs *corev1.SecretList
	var err error

	if len(labels) == 0 {
		secs, err = s.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		secs, err = s.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return secs.Items, nil
}
