package configmaps

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/util/retry"
)

type ConfigMapClient struct {
	client v1.ConfigMapInterface
	ns     string
}

func NewConfigMapClient(clientset *kubernetes.Clientset, namespace string) *ConfigMapClient {
	return &ConfigMapClient{
		client: clientset.CoreV1().ConfigMaps(namespace),
		ns:     namespace,
	}
}

func (c *ConfigMapClient) Create(name string, data map[string]string) (*corev1.ConfigMap, error) {
	cm := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: c.ns,
		},
		Data: data,
	}
	return c.client.Create(context.TODO(), cm, metav1.CreateOptions{})
}

func (c *ConfigMapClient) Update(name string, data map[string]string) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		cm, getErr := c.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		cm.Data = data

		_, updateErr := c.client.Update(context.TODO(), cm, metav1.UpdateOptions{})
		return updateErr
	})
	return retryErr
}

func (c *ConfigMapClient) Delete(name string) error {
	return c.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (c *ConfigMapClient) Get(name string) (*corev1.ConfigMap, error) {
	return c.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *ConfigMapClient) List(labels ...string) ([]corev1.ConfigMap, error) {
	var cms *corev1.ConfigMapList
	var err error

	if len(labels) == 0 {
		cms, err = c.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		cms, err = c.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return cms.Items, nil
}
