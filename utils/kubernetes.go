package utils

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func ConnectToKubernets() {
	// Create a new Kubernetes client
	clientset, err := GetClientSet()

	if err != nil {
		log.Fatal(err)
	}

	// List all pods in the default namespace
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d pods in the default namespace\n", len(pods.Items))
}

func GetClientSet() (*kubernetes.Clientset, error) {
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	return kubernetes.NewForConfig(config)

}

func GetConfig() (*rest.Config, error) {
	const configPath = "/etc/rancher/k3s/k3s.yaml"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	return config, err
}

func CreateNamespace(name string) error {
	clientset, err := GetClientSet()
	if err != nil {
		return err
	}
	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	return err
}
