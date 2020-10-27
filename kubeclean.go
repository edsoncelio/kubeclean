/*
Tool to clean empty namespaces on kubernetes!
*/

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string

	//protected namespaces
	namespaceExceptions := [4]string{
		"default",
		"kube-system",
		"kube-node-lease",
		"kube-public",
	}

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

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get namespaces:", err)
	}

	fmt.Printf("kubeclean - remove empty namespaces \n")

	// print namespaces
	for _, namespace := range namespaces.Items {

		result := checkEx(namespaceExceptions, namespace.GetName())
		if !result {
			fmt.Printf("Deployments in namespace %s\n", namespace.GetName())

			//list deployments
			deployments, err := clientset.AppsV1().Deployments(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get deployments: ", err)
			}

			//list statefulsets
			statefulsets, err := clientset.AppsV1().StatefulSets(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get statefulset: ", err)
			}

			deployCount := len(deployments.Items)
			statefulsetCount := len(statefulsets.Items)

			//check if exists one or more deployments or statefulset
			if deployCount < 1 {
				fmt.Println("Deploy not found, clear the namespace?")
				prompt()
				err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace.GetName(), metav1.DeleteOptions{})
				if err != nil {
					log.Fatalln("failed to delete namespace: ", err)
				} else {
					log.Fatalln("Namespace deleted!: ")
				}
			} else if statefulsetCount < 1 {
				fmt.Println("Deploy not found, clear the namespace?")
				prompt()
				err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace.GetName(), metav1.DeleteOptions{})
				if err != nil {
					log.Fatalln("failed to delete namespace: ", err)
				} else {
					log.Fatalln("Namespace deleted!: ")
				}
			} else {
				//print deployments
				for _, deployment := range deployments.Items {
					fmt.Printf("%s\n", deployment.GetName())
				}
			}

			fmt.Printf(" ------------------ \n")

		} else {
			fmt.Printf(" Protected namespace: %s \n", namespace.GetName())
		}
	}

}

func checkEx(arr [4]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
