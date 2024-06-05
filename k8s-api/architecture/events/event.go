package events

import (
	"context"

	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/events/v1"
)

type EventClient struct {
	client v1.EventInterface
	ns     string
}

func NewEventClient(clientset *kubernetes.Clientset, namespace string) *EventClient {
	return &EventClient{
		client: clientset.EventsV1().Events(namespace),
		ns:     namespace,
	}
}

func (c *EventClient) Create(name string) (*eventsv1.Event, error) {
	e := &eventsv1.Event{}
	return c.client.Create(context.TODO(), e, metav1.CreateOptions{})
}

func (c *EventClient) Delete(name string) error {
	return c.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (c *EventClient) Get(name string) (*eventsv1.Event, error) {
	return c.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *EventClient) List(labels ...string) ([]eventsv1.Event, error) {
	var el *eventsv1.EventList
	var err error

	if len(labels) == 0 {
		el, err = c.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		el, err = c.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}
	if err != nil {
		return nil, err
	}

	return el.Items, nil
}
