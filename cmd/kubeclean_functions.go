/*
Tool to clean empty namespaces on kubernetes!
*/

package cmd

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/manifoldco/promptui"
)

func checkNamespaceDelete(clientset *kubernetes.Clientset, namespaceDelete []string, result string) {
	switch {
	case result == "all":
		for _, namespace := range namespaceDelete {
			err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
			if err != nil {
				log.Fatalln("failed to delete namespace: ", err)
			} else {
				fmt.Printf("🔥  Namespace %s deleted!\n", namespace)
			}
		}
		fmt.Printf("\n✅ Done!\n")
	case result == "exit":
		fmt.Printf("\n☑️  Cancelled!\n")
	case result != "all" && result != "exit":
		err := clientset.CoreV1().Namespaces().Delete(context.TODO(), result, metav1.DeleteOptions{})
		if err != nil {
			log.Fatalln("failed to delete namespace: ", err)
		} else {
			fmt.Printf("🔥  Namespace %s deleted!", result)
		}
		fmt.Printf("\n✅ Done!\n")

	}
}

func execNamespaceCheck(kubeconfig string) {

	//namespaces to delete
	var namespaceDelete []string

	//protected namespaces
	namespaceExceptions := []string{"default", "kube-system", "kube-public"}

	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }

	// flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
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

	for _, namespace := range namespaces.Items {

		result := checkEx(namespaceExceptions, namespace.GetName())
		if !result {

			//deployments
			deployments, err := clientset.AppsV1().Deployments(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get deployments: ", err)
			}

			//statefulsets
			statefulsets, err := clientset.AppsV1().StatefulSets(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get statefulset: ", err)
			}

			//services
			services, err := clientset.CoreV1().Services(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get service: ", err)
			}

			//serviceaccounts
			serviceaccounts, err := clientset.CoreV1().ServiceAccounts(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get serviceaccount: ", err)
			}

			//secrets
			secrets, err := clientset.CoreV1().Secrets(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get secrets: ", err)
			}

			deployCount := len(deployments.Items)
			statefulsetCount := len(statefulsets.Items)
			serviceCount := len(services.Items)
			serviceaccountCount := len(serviceaccounts.Items)
			secretCount := len(secrets.Items)

			switch {
			case deployCount < 1:
				namespaceDelete = append(namespaceDelete, namespace.GetName())
			case deployCount < 1 && serviceCount == 1 && secretCount == 1:
				namespaceDelete = append(namespaceDelete, namespace.GetName())
			case statefulsetCount < 1 && serviceaccountCount == 1 && secretCount == 1 && deployCount < 1:
				namespaceDelete = append(namespaceDelete, namespace.GetName())
			default:
				for _, deployment := range deployments.Items {
					fmt.Printf("\n🤖 Found namespace '%s' with deployments: %s \n", namespace.GetName(), deployment.GetName())
				}
			}

		} else {
			fmt.Printf("❗ Found a protected namespace: %s ⏩ \n", namespace.GetName())
		}
	}
	if len(namespaceDelete) > 0 {

		result := yesNo(namespaceDelete)
		checkNamespaceDelete(clientset, namespaceDelete, result)

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
	namespaceDelete = append(namespaceDelete, "exit")

	prompt := promptui.Select{
		Label: "Namespaces to remove:",
		Items: namespaceDelete,
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}
