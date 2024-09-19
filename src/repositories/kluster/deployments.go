package kluster

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListDeployments() ([]v1.Deployment, error) {
	clientset, err := getConfig()
	if err != nil {
		return nil, err
	}

	deploymentList, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	deployments := make([]v1.Deployment, 0)
	for _, deployment := range deploymentList.Items {
		annotations := deployment.GetAnnotations()
		if annotations == nil {
			continue
		}

		if annotations["scaler.reyel.cloud/enabled"] == "true" {
			deployments = append(deployments, deployment)
		}
	}

	return deployments, nil
}

func ScaleDeployment(deployment v1.Deployment, replicas int32) error {
	clientset, err := getConfig()
	if err != nil {
		return err
	}

	deployment.Spec.Replicas = &replicas
	_, err = clientset.AppsV1().Deployments(deployment.Namespace).Update(context.TODO(), &deployment, metav1.UpdateOptions{})

	return err
}
