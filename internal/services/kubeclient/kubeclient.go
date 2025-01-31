package kubeclient

import (
	"context"
	"fmt"

	"fallback/internal/config"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KubeClient struct {
	clientset *kubernetes.Clientset
}

func NewKubeClient(cfg config.ServerConfigs) (*KubeClient, error) {
	var config *rest.Config
	var err error

	if cfg.UseServiceAccount {
		config, err = rest.InClusterConfig()
	} else {
		kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	}

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubeClient{clientset}, nil
}

type ServiceReference struct {
	name      string
	namespace string
}

var emptyServiceRef = ServiceReference{}

func (kubeclient *KubeClient) getServiceRefForHostname(hostname string) (ServiceReference, error) {
	ingressList, err := kubeclient.clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return emptyServiceRef, err
	}

	for _, ingress := range ingressList.Items {
		for _, rule := range ingress.Spec.Rules {
			if rule.Host == hostname {
				if len(rule.HTTP.Paths) != 1 {
					return emptyServiceRef, fmt.Errorf("non unique path for host")
				}

				return ServiceReference{
					name:      rule.HTTP.Paths[0].Backend.Service.Name,
					namespace: ingress.Namespace,
				}, nil
			}
		}
	}

	return emptyServiceRef, fmt.Errorf("rule not found for host")
}

type DeploymentReference struct {
	name      string
	namespace string
}

var emptyDeploymentRef = DeploymentReference{}

func (kubeclient *KubeClient) getDeployRefForService(serviceRef ServiceReference) (DeploymentReference, error) {
	service, err := kubeclient.clientset.CoreV1().Services(serviceRef.namespace).Get(context.TODO(), serviceRef.name, metav1.GetOptions{})
	if err != nil {
		return emptyDeploymentRef, err
	}

	return DeploymentReference{
		name:      service.Spec.Selector["app"],
		namespace: service.Namespace,
	}, nil
}

func (kubeclient *KubeClient) getDeployment(ref DeploymentReference) (*v1.Deployment, error) {
	return kubeclient.clientset.AppsV1().Deployments(ref.namespace).Get(context.TODO(), ref.name, metav1.GetOptions{})
}

func (kubeclient *KubeClient) GetDeploymentForHostname(hostname string) (*appv1.Deployment, error) {
	serviceRef, err := kubeclient.getServiceRefForHostname(hostname)
	if err != nil {
		return nil, err
	}

	deployRef, err := kubeclient.getDeployRefForService(serviceRef)
	if err != nil {
		return nil, err
	}

	return kubeclient.getDeployment(deployRef)
}

func (kubeclient *KubeClient) ScaleDeployment(deployment *v1.Deployment, replicas int32) error {
	deployment.Spec.Replicas = &replicas

	_, err := kubeclient.clientset.AppsV1().Deployments(deployment.Namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})

	return err
}
