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

	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deployments.Items, nil
}
