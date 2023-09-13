package wrappers

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"xs/pkg/io"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/exec"
	"strings"
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
		io.Fatal(err)
	}

	// List all pods in the default namespace
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		io.Fatal(err)
	}
	io.Printf("There are %d pods in the default namespace\n", len(pods.Items))
}

func (k *Kuber) GetClientSet() (*kubernetes.Clientset, error) {
	config, err := k.GetConfig()
	if err != nil {
		io.Fatal(err)
	}
	return kubernetes.NewForConfig(config)

}
func (k *Kuber) ExportToFile(file string) error {
	configPath := os.Getenv(K3S_ENV)

	command := "kubectl config view --raw > " + file

	split := strings.Split(command, " ")

	io.Println("K3S_ENV: ", configPath)

	err := exec.Command(split[0], split[1:]...).Run()
	return err
}

func getTempDir() (string, error) {
	return os.MkdirTemp("", "k3s_conf")
}

func (k *Kuber) GetConfig() (*rest.Config, error) {
	configPath := os.Getenv(K3S_ENV)

	io.Println("K3S_ENV: ", configPath)

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

func (k *Kuber) CreateServiceAccount(name string, ns string) error {
	serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
	}

	// Create the service account in the cluster.
	_, err := k.clientset.CoreV1().ServiceAccounts(ns).Create(context.Background(), serviceAccount, metav1.CreateOptions{})
	if err != nil {
		io.Printf("Error creating service account: %v\n", err)
	}

	io.Println("Service account created successfully.")
	return err
}
