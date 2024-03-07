package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	content, err := os.ReadFile("config")
	if err != nil {
		fmt.Printf("Can not read kubeconfig file: %v\n", err)
	}
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(content))
	if err != nil {
		fmt.Printf("Can not read kubeconfig file: %v\n", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Can not read kubeconfig file: %v\n", err)
	}

	pod, err := clientset.CoreV1().Pods("kube-system").Get(context.TODO(), "coredns-6799fbcd5-flblk", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Can not read pods: %v\n", err)
	}
	fmt.Print(pod, "\n\n\n\n\n")

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	fmt.Print(nodes)
}
