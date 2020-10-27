/*
Tool to clean empty namespaces on kubernetes!
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/manifoldco/promptui"
)

func main() {
	var kubeconfig *string

	//namespaces to delete
	var namespaceDelete []string

	//protected namespaces
	var namespaceExceptions []string

	namespaceExceptions = append(namespaceExceptions, "default")
	namespaceExceptions = append(namespaceExceptions, "kube-system")
	namespaceExceptions = append(namespaceExceptions, "kube-node-lease")
	namespaceExceptions = append(namespaceExceptions, "kube-public")

	//delete policies
	//var deletePolicies []string

	/*
		deletePolicies := [2]string{
			"Deployment",
			"Service",
		}

	*/

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

	fmt.Printf("⚠️  kubeclean - remove empty namespaces ⚠️ \n\n")
	fmt.Printf("🔍 Searching resources... \n\n")

	// print namespaces
	for _, namespace := range namespaces.Items {

		result := checkEx(namespaceExceptions, namespace.GetName())
		if !result {
			//fmt.Printf("Deployments in namespace %s\n", namespace.GetName())

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
				namespaceDelete = append(namespaceDelete, namespace.GetName())
				/*
					                err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace.GetName(), metav1.DeleteOptions{})
									if err != nil {
										log.Fatalln("failed to delete namespace: ", err)
									} else {
										log.Fatalln("Namespace deleted!: ")
					                }
				*/
			} else if statefulsetCount < 1 {
				fmt.Println("Deploy not found, clear the namespace?")
				namespaceDelete = append(namespaceDelete, namespace.GetName())
				/*
									err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace.GetName(), metav1.DeleteOptions{})
									if err != nil {
										log.Fatalln("failed to delete namespace: ", err)
									} else {
										log.Fatalln("Namespace deleted!: ")
					                }
				*/
			} else {
				//print deployments
				for _, deployment := range deployments.Items {
					fmt.Printf("%s\n", deployment.GetName())
				}
			}

			//fmt.Printf(" ------------------ \n")

		} else {
			fmt.Printf("❗ Found a protected namespace: %s ⏩ \n", namespace.GetName())
		}
	}
	if len(namespaceDelete) > 0 {

		result := yesNo(namespaceDelete)

		if result != "all" {
			err := clientset.CoreV1().Namespaces().Delete(context.TODO(), result, metav1.DeleteOptions{})
			if err != nil {
				log.Fatalln("failed to delete namespace: ", err)
			} else {
				fmt.Printf("🔥  Namespace %s deleted!", result)
			}
			fmt.Printf("\n✅ Done!\n")
		} else if result == "all" {
			for _, namespace := range namespaceDelete {
				err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
				if err != nil {
					log.Fatalln("failed to delete namespace: ", err)
				} else {
					fmt.Printf("🔥  Namespace %s deleted!\n", namespace)
				}
			}
			fmt.Printf("\n✅ Done!\n")
		} else {
			fmt.Printf("\n☑️  Cancelled!\n")
		}

	} else {
		fmt.Printf("\nNothing to remove, congrats! 🎉\n")
	}

}

func checkEx(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func yesNo(namespaceDelete []string) string {

	namespaceDelete = append(namespaceDelete, "all")

	prompt := promptui.Select{
		Label: "Remove the namespaces?",
		Items: namespaceDelete,
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}
