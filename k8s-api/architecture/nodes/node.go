package nodes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/util/retry"
)

type NodeClient struct {
	client v1.NodeInterface
}

func NewNodeClient(clientset *kubernetes.Clientset) *NodeClient {
	return &NodeClient{
		client: clientset.CoreV1().Nodes(),
	}
}

func (n *NodeClient) Create(name string) (*corev1.Node, error) {
	nod := &corev1.Node{}
	return n.client.Create(context.TODO(), nod, metav1.CreateOptions{})
}

func (n *NodeClient) Update(name string) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		nod, getErr := n.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		// nod.

		_, updateErr := n.client.Update(context.TODO(), nod, metav1.UpdateOptions{})
		return updateErr
	})
	return retryErr
}

func (n *NodeClient) Delete(name string) error {
	return n.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (n *NodeClient) Get(name string) (*corev1.Node, error) {
	return n.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (n *NodeClient) List(labels ...string) ([]corev1.Node, error) {
	var nods *corev1.NodeList
	var err error

	if len(labels) == 0 {
		nods, err = n.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		nods, err = n.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}
	if err != nil {
		return nil, err
	}

	return nods.Items, nil
}
