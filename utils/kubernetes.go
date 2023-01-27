package utils

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func ConnectToKubernets() {
	// Create a new Kubernetes client
	const configPath = "/etc/rancher/k3s/k3s.yaml"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// List all pods in the default namespace
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d pods in the default namespace\n", len(pods.Items))
}
