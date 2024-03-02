package v1

import (
	"context"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	"k8s.io/client-go/util/retry"
)

type HPAClient struct {
	client v1.HorizontalPodAutoscalerInterface
	ns     string
}

func NewHPAClient(clientset *kubernetes.Clientset, namespace string) *HPAClient {
	return &HPAClient{
		client: clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace),
		ns:     namespace,
	}
}

func (h *HPAClient) Create(name string, minReplicas, maxReplicas, cpuUtilization int32, deploymentName string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	hpa := &autoscalingv1.HorizontalPodAutoscaler{
		TypeMeta: metav1.TypeMeta{
			Kind:       "HorizontalPodAutoscaler",
			APIVersion: "autoscaling/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: h.ns,
		},
		Spec: autoscalingv1.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: autoscalingv1.CrossVersionObjectReference{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
				Name:       deploymentName,
			},
			MinReplicas:                    &minReplicas,
			MaxReplicas:                    maxReplicas,
			TargetCPUUtilizationPercentage: &cpuUtilization,
		},
	}
	return h.client.Create(context.TODO(), hpa, metav1.CreateOptions{})
}

func (h *HPAClient) Update(name string, minReplicas, maxReplicas, cpuUtilization int32) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		hpa, getErr := h.client.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		hpa.Spec.MinReplicas = &minReplicas
		hpa.Spec.MaxReplicas = maxReplicas
		hpa.Spec.TargetCPUUtilizationPercentage = &cpuUtilization

		_, updateErr := h.client.Update(context.TODO(), hpa, metav1.UpdateOptions{})
		return updateErr
	})

	return retryErr
}

func (h *HPAClient) Delete(name string) error {
	return h.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (h *HPAClient) Get(name string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	return h.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (h *HPAClient) List(labels ...string) ([]autoscalingv1.HorizontalPodAutoscaler, error) {
	var hpaList *autoscalingv1.HorizontalPodAutoscalerList
	var err error

	if len(labels) == 0 {
		hpaList, err = h.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		hpaList, err = h.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return hpaList.Items, nil
}
