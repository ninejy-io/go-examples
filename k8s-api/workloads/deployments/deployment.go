package deployments

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type DeploymentClient struct {
	client v1.DeploymentInterface
	ns     string
}

func NewDeploymentClient(clientset *kubernetes.Clientset, namespace string) *DeploymentClient {
	return &DeploymentClient{
		client: clientset.AppsV1().Deployments(namespace),
		ns:     namespace,
	}
}

func (d *DeploymentClient) Create(name string, labels map[string]string, replicas int32, image string) {
	var containers []corev1.Container
	container := corev1.Container{
		Name:  name,
		Image: image,
		// Command: []string,
		// Args: []string,
		Ports: []corev1.ContainerPort{},
		// EnvFrom:   []corev1.EnvFromSource{},
		Env:       []corev1.EnvVar{},
		Resources: corev1.ResourceRequirements{},
		// RestartPolicy: ,
		// VolumeMounts:   []corev1.VolumeMount{},
		LivenessProbe:  &corev1.Probe{},
		ReadinessProbe: &corev1.Probe{},
		StartupProbe:   &corev1.Probe{},
		// Lifecycle: ,
		ImagePullPolicy: corev1.PullIfNotPresent,
	}
	containers = append(containers, container)

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: d.ns,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   name,
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					// Volumes: ,
					// InitContainers: ,
					Containers: containers,
					// RestartPolicy: ,
					// TerminationGracePeriodSeconds: ,
					// ActiveDeadlineSeconds: ,
					// DNSPolicy: ,
					// NodeSelector: ,
					// ServiceAccountName: ,
					// ImagePullSecrets: ,
					// Affinity: ,
					// Tolerations: ,
					// ReadinessGates: ,
				},
			},
		},
	}

	d.client.Create(context.TODO(), deployment, metav1.CreateOptions{})
}

func (d *DeploymentClient) Update(name string) {}

func (d *DeploymentClient) Delete(name string) error {
	return d.client.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (d *DeploymentClient) Get(name string) (*appsv1.Deployment, error) {
	return d.client.Get(context.TODO(), name, metav1.GetOptions{})
}

func (d *DeploymentClient) List(labels ...string) ([]appsv1.Deployment, error) {
	var deployments *appsv1.DeploymentList
	var err error

	if len(labels) == 0 {
		deployments, err = d.client.List(context.TODO(), metav1.ListOptions{})
	} else {
		deployments, err = d.client.List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels[0],
		})
	}

	if err != nil {
		return nil, err
	}
	return deployments.Items, nil
}
