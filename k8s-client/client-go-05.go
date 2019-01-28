package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main()  {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	name := "client-go-test"
	namespacesClient := clientset.CoreV1().Namespaces()
	namespace := &apiv1.Namespace{
		ObjectMeta : metav1.ObjectMeta {
			Name: name,
		},
		Status: apiv1.NamespaceStatus{
			Phase: apiv1.NamespaceActive,
		},
	}

	fmt.Println("Creating Namespaces...")
	result, err := namespacesClient.Create(namespace)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created Namespaces %s on %s\n", result.ObjectMeta.Name, result.ObjectMeta.CreationTimestamp)


	fmt.Println("Creating Namespaces...")
	result, err = namespacesClient.Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s, Status: %s, selfLink: %s, uid: %s\n", result.ObjectMeta.Name, result.Status.Phase, result.ObjectMeta.SelfLink, result.ObjectMeta.UID)

	fmt.Println("Deleting Namespaces...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := namespacesClient.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Printf("Deleted Namespaces %s\n", name)
}
