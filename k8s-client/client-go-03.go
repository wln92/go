package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	var kubeconfig *string

	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	for {
		//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		//
		//if err != nil {
		//	panic(err.Error())
		//}
		//
		//fmt.Printf("There are %d pods in the k8s cluster\n", len(pods.Items))


		namespace := "ingress-nginx"
		pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nThere are %d pods in the namespaces %s\n", len(pods.Items), namespace)
		for _, pod := range pods.Items {
			fmt.Printf("Name:%s, Status:%s, CreateTime:%s\n", pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
		}

		podName := "nginx-ingress-controller-m87qg"
		pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", podName, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n", podName, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("\nFound pod %s in namespace %s\n", podName, namespace)
			maps := map[string]interface{} {
				"Name": pod.ObjectMeta.Name,
				"Namespaces": pod.ObjectMeta.Namespace,
				"NodeName": pod.Spec.NodeName,
				"Annotations": pod.ObjectMeta.Annotations,
				"Labels": pod.ObjectMeta.Labels,
				"SelfLink": pod.ObjectMeta.SelfLink,
				"Uid": pod.ObjectMeta.UID,
				"Status": pod.Status.Phase,
				"IP": pod.Status.PodIP,
				"Image": pod.Spec.Containers[0].Image,
			}
			prettyPrint(maps)
		}

		logs := getPodLogs(podName, namespace, clientset)

		fmt.Printf("podName:%s in %s logs: %s\n", podName, namespace, logs)
		time.Sleep(60 * time.Second)



	}
}

func getPodLogs(podName, namespace string, clientset *kubernetes.Clientset) string {
	podLogOpts := v1.PodLogOptions{}
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream()
	if err != nil {
		return "error in opening stream"
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "error in copy information from podLogs to buf"
	}
	str := buf.String()

	return str
}

func prettyPrint(maps map[string]interface{}) {
	lens := 0
	for k, _ := range maps {
		if lens <= len(k) {
			lens = len(k)
		}
	}
	for key, values := range maps {
		spaces := lens - len(key)
		v := ""
		for i := 0; i < spaces; i++ {
			v += " "
		}
		fmt.Printf("%s: %s%v\n", key, v, values)
	}

}


func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}