package service

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/util/retry"
)

type ServiceClient struct {
	client v1.ServiceInterface
	ns     string
}

func NewServiceClient(clientset *kubernetes.Clientset, namespace string) *ServiceClient {
	return &ServiceClient{
		client: clientset.CoreV1().Services(namespace),
		ns:     namespace,
	}
}

func (s *ServiceClient) Create(name string, selector map[string]string, svcType corev1.ServiceType, portName string, port, targetPort int32, nodePort ...int32) (*corev1.Service, error) {
	var ports []corev1.ServicePort
	_port := corev1.ServicePort{
		Name:     portName,
		Protocol: "TCP",
		Port:     port,
		TargetPort: intstr.IntOrString{
			IntVal: targetPort,
		},
	}
	if len(nodePort) != 0 {
		_port.NodePort = nodePort[0]
	}
	ports = append(ports, _port)

	svc := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: s.ns,
		},
		Spec: corev1.ServiceSpec{
			Ports:    ports,
			Selector: selector,
			Type:     svcType,
		},
	}
	return s.client.Create(context.TODO(), svc, metav1.CreateOptions{})
}

func (s *ServiceClient) Update(name string, selector map[string]string, svcType corev1.ServiceType, portName string, port, targetPort int32, nodePort ...int32) error {
	var ports []corev1.ServicePort
	_port := corev1.ServicePort{
		Name:     portName,
		Protocol: "TCP",
		Port:     port,
		TargetPort: intstr.IntOrString{
			IntVal: targetPort,
		},
	}
	if len(nodePort) != 0 {
		_port.NodePort = nodePort[0]
	}
	ports = append(ports, _port)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		svc, getErr := s.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		svc.Spec.Ports = ports
		svc.Spec.Selector = selector
		svc.Spec.Type = svcType

		_, updateErr := s.client.Update(context.TODO(), svc, metav1.UpdateOptions{})
		return updateErr
	})

	return retryErr
}

func (s *ServiceClient) Delete(name string) error {
	return s.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (s *ServiceClient) Get(name string) (*corev1.Service, error) {
	return s.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (s *ServiceClient) List(labels ...string) ([]corev1.Service, error) {
	var services *corev1.ServiceList
	var err error

	if len(labels) == 0 {
		services, err = s.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		services, err = s.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return services.Items, nil
}
