package ingress

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/util/retry"
)

type IngressClient struct {
	client v1.IngressInterface
	ns     string
}

func NewIngressClient(clientset *kubernetes.Clientset, namespace string) *IngressClient {
	return &IngressClient{
		client: clientset.NetworkingV1().Ingresses(namespace),
		ns:     namespace,
	}
}

type SimplePathConfig struct {
	Path        string
	PathType    *networkingv1.PathType
	ServiceName string
	ServicePort int32
}

func (i *IngressClient) Create(name string, ingressClassName string, host string, simplePaths []SimplePathConfig) (*networkingv1.Ingress, error) {
	var rules []networkingv1.IngressRule
	var paths []networkingv1.HTTPIngressPath

	for _, item := range simplePaths {
		_path := networkingv1.HTTPIngressPath{
			Path:     item.Path,
			PathType: item.PathType,
			Backend: networkingv1.IngressBackend{
				Service: &networkingv1.IngressServiceBackend{
					Name: item.ServiceName,
					Port: networkingv1.ServiceBackendPort{
						Number: item.ServicePort,
					},
				},
			},
		}
		paths = append(paths, _path)
	}

	rule := networkingv1.IngressRule{
		Host: host,
		IngressRuleValue: networkingv1.IngressRuleValue{
			HTTP: &networkingv1.HTTPIngressRuleValue{
				Paths: paths,
			},
		},
	}
	rules = append(rules, rule)

	ing := &networkingv1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.ns,
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: &ingressClassName,
			Rules:            rules,
		},
	}
	return i.client.Create(context.TODO(), ing, metav1.CreateOptions{})
}

func (i *IngressClient) Update(name string, ingressClassName string, host string, simplePaths []SimplePathConfig) error {
	var paths []networkingv1.HTTPIngressPath

	for _, item := range simplePaths {
		_path := networkingv1.HTTPIngressPath{
			Path:     item.Path,
			PathType: item.PathType,
			Backend: networkingv1.IngressBackend{
				Service: &networkingv1.IngressServiceBackend{
					Name: item.ServiceName,
					Port: networkingv1.ServiceBackendPort{
						Number: item.ServicePort,
					},
				},
			},
		}
		paths = append(paths, _path)
	}

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		ing, getErr := i.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		oldRules := ing.Spec.Rules

		tempMap := make(map[string]struct{}, len(oldRules))
		for _, oldRule := range oldRules {
			tempMap[oldRule.Host] = struct{}{}
		}
		if _, ok := tempMap[host]; ok {
			// 如果 host 已经存在, 在该 host 下添加 path 即可
			for _, item := range ing.Spec.Rules {
				if item.Host == host {
					item.IngressRuleValue.HTTP.Paths = append(item.IngressRuleValue.HTTP.Paths, paths...)
				}
			}
		} else {
			// 如果 host 不存在, 添加 host 和 path
			rule := networkingv1.IngressRule{
				Host: host,
				IngressRuleValue: networkingv1.IngressRuleValue{
					HTTP: &networkingv1.HTTPIngressRuleValue{
						Paths: paths,
					},
				},
			}
			ing.Spec.Rules = append(ing.Spec.Rules, rule)
		}

		ing.Spec.IngressClassName = &ingressClassName

		_, updateErr := i.client.Update(context.TODO(), ing, metav1.UpdateOptions{})
		return updateErr
	})
	return retryErr
}

func (i *IngressClient) Delete(name string) error {
	return i.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (i *IngressClient) Get(name string) (*networkingv1.Ingress, error) {
	return i.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (i *IngressClient) List(labels ...string) ([]networkingv1.Ingress, error) {
	var ingresses *networkingv1.IngressList
	var err error

	if len(labels) == 0 {
		ingresses, err = i.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		ingresses, err = i.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return ingresses.Items, err
}
