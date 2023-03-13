package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// create the Kubernetes API client
	fmt.Println(os.Environ())
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get the pod name and namespace
	podName := os.Getenv("HOSTNAME")
	fmt.Println("podName: ", podName)
	// get the pod object
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// get the node IP address
	nodeIP := pod.Status.HostIP
	fmt.Println("Node IP address:", nodeIP)
}
