package main

import (
	"fmt"
	k8sapi "k8s-api"
	"k8s-api/services/service"
)

func main() {
	clientSet := k8sapi.InitClientset()

	// podClient := pods.NewPodClient(clientSet, "default")
	// allPods, _ := podClient.List()
	// fmt.Println(allPods)
	// allPods, _ := podClient.List("app=xxl-job-admin")
	// fmt.Println(allPods)
	// pod, _ := podClient.Get("xxl-job-admin-95f698754-7h4f2")
	// fmt.Printf("%v\n", pod)

	// hpaClient := autoscalingv1.NewHPAClient(clientSet, "default")
	// fmt.Println(hpaClient.List())
	// fmt.Println(hpaClient.Get("hpa-nginx-deployment"))
	// _ = hpaClient.Update("hpa-nginx-deployment", 1, 4, 60)

	// hpaClientv2 := autoscalingv2.NewHPAClient(clientSet, "default")
	// hpaClientv2.Update("hpa-nginx-deployment", 1, 2, 80)
	// hpaClientv2.Delete("hpa-nginx-deployment")
	// hpaClientv2.Create("hpa-nginx-deployment", 1, 2, 80, "nginx-deployment")

	svcClient := service.NewServiceClient(clientSet, "default")
	// fmt.Println(svcClient.Get("deployment-svc-nginx"))
	// svcClient.Delete("deployment-svc-nginx")
	fmt.Println(svcClient.Create("deployment-svc-nginx", map[string]string{"app": "nginx"}, "ClusterIP", "web", 80, 80))
}
