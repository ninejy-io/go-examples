package v2

import (
	"context"
	"errors"

	autoscalingv2 "k8s.io/api/autoscaling/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	"k8s.io/client-go/util/retry"
)

type HPAClient struct {
	client v2.HorizontalPodAutoscalerInterface
	ns     string
}

func NewHPAClient(clientset *kubernetes.Clientset, namespace string) *HPAClient {
	return &HPAClient{
		client: clientset.AutoscalingV2().HorizontalPodAutoscalers(namespace),
		ns:     namespace,
	}
}

func (h *HPAClient) Create(name string, minReplicas, maxReplicas, cpuUtilization, memoryUtilization int32, deploymentName string) (*autoscalingv2.HorizontalPodAutoscaler, error) {
	var metrics []autoscalingv2.MetricSpec
	// 如果不依靠CPU使用率进行伸缩, cpuUtilization = 0
	if cpuUtilization != 0 {
		cpuMetric := autoscalingv2.MetricSpec{
			Type: "Resource",
			Resource: &autoscalingv2.ResourceMetricSource{
				Name: "cpu",
				Target: autoscalingv2.MetricTarget{
					Type:               "Utilization",
					AverageUtilization: &cpuUtilization,
				},
			},
		}
		metrics = append(metrics, cpuMetric)
	}

	// 如果不依靠memory使用率进行伸缩, memoryUtilization = 0
	if memoryUtilization != 0 {
		memoryMetric := autoscalingv2.MetricSpec{
			Type: "Resource",
			Resource: &autoscalingv2.ResourceMetricSource{
				Name: "memory",
				Target: autoscalingv2.MetricTarget{
					Type:               "Utilization",
					AverageUtilization: &memoryUtilization,
				},
			},
		}
		metrics = append(metrics, memoryMetric)
	}

	if metrics == nil {
		return nil, errors.New("no resource utilization provided")
	}

	hpa := &autoscalingv2.HorizontalPodAutoscaler{
		TypeMeta: metav1.TypeMeta{
			Kind:       "HorizontalPodAutoscaler",
			APIVersion: "autoscaling/v2",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: h.ns,
		},
		Spec: autoscalingv2.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: autoscalingv2.CrossVersionObjectReference{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
				Name:       deploymentName,
			},
			MinReplicas: &minReplicas,
			MaxReplicas: maxReplicas,
			Metrics:     metrics,
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
		// fmt.Printf("hpa.Spec.Metrics: %v\n", hpa)
		// hpa.Spec.Metrics[0].Resource.Target.AverageUtilization = &cpuUtilization

		var metrics []autoscalingv2.MetricSpec
		cpuMetric := autoscalingv2.MetricSpec{
			Type: "Resource",
			Resource: &autoscalingv2.ResourceMetricSource{
				Name: "cpu",
				Target: autoscalingv2.MetricTarget{
					Type:               "Utilization",
					AverageUtilization: &cpuUtilization,
				},
			},
		}
		metrics = append(metrics, cpuMetric)

		hpa.Spec.Metrics = metrics

		_, updateErr := h.client.Update(context.TODO(), hpa, metav1.UpdateOptions{})
		return updateErr
	})

	return retryErr
}

func (h *HPAClient) Delete(name string) error {
	return h.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (h *HPAClient) Get(name string) (*autoscalingv2.HorizontalPodAutoscaler, error) {
	return h.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (h *HPAClient) List(labels ...string) ([]autoscalingv2.HorizontalPodAutoscaler, error) {
	var hpaList *autoscalingv2.HorizontalPodAutoscalerList
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
