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
	"os"
)

const K3S_ENV = "K3S_CONF"
const LINUX_DEFAULT_K3S_CONF = "/etc/rancher/k3s/k3s.yaml"

type Kuber struct {
	clientset *kubernetes.Clientset
}

func (k *Kuber) ConnectToKubernetes() {
	// Create a new Kubernetes client
	clientset, err := k.GetClientSet()

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

func (k *Kuber) GetClientSet() (*kubernetes.Clientset, error) {
	config, err := k.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	return kubernetes.NewForConfig(config)

}

func (k *Kuber) GetConfig() (*rest.Config, error) {
	configPath := os.Getenv(K3S_ENV)

	println("K3S_ENV: ", configPath)

	if configPath == "" {
		configPath = LINUX_DEFAULT_K3S_CONF
	}

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	return config, err
}

func (k *Kuber) CreateNamespace(name string) error {
	clientset, err := k.GetClientSet()
	if err != nil {
		return err
	}
	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	return err
}
