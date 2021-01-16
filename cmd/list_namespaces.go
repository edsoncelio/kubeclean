package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var runCmd = &cobra.Command{
	Use:              "list-all",
	Short:            "List namespaces",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		runAll(cmd)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	if home := homedir.HomeDir(); home != "" {
		runCmd.PersistentFlags().StringP("config", "", filepath.Join(home, ".kube", "config"), "kubectl config path")
	} else {
		runCmd.PersistentFlags().StringP("config", "", "", "kubectl config path")

	}
	runCmd.PersistentFlags().Bool("remove", false, "flag to list and remove, default: false")
}

func checkEx(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func runAll(cmd *cobra.Command) {
	//namespaces to delete
	var namespaceDelete []string

	//protected namespaces
	namespaceExceptions := []string{"default", "kube-system", "kube-public"}

	kubeConfig, err := cmd.Flags().GetString("config")
	if err != nil {
		fmt.Print("Error to parse kubeconfig.")
		os.Exit(1)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
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

	fmt.Printf("🔍 Searching resources... \n\n")

	// print namespaces
	for _, namespace := range namespaces.Items {

		result := checkEx(namespaceExceptions, namespace.GetName())
		// fmt.Println(namespace.GetName())
		if !result {
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

			//list services
			services, err := clientset.CoreV1().Services(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get service: ", err)
			}

			//list serviceaccounts
			serviceaccounts, err := clientset.CoreV1().ServiceAccounts(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get serviceaccount: ", err)
			}

			//list secrets
			secrets, err := clientset.CoreV1().Secrets(namespace.GetName()).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln("failed to get secrets: ", err)
			}

			deployCount := len(deployments.Items)
			statefulsetCount := len(statefulsets.Items)
			serviceCount := len(services.Items)
			serviceaccountCount := len(serviceaccounts.Items)
			secretCount := len(secrets.Items)

			//check if exists one or more deployments or statefulset
			switch {
			case deployCount < 1:
				namespaceDelete = append(namespaceDelete, namespace.GetName())
			case deployCount < 1 && serviceCount == 1 && secretCount == 1:
				namespaceDelete = append(namespaceDelete, namespace.GetName())
			case statefulsetCount < 1 && serviceaccountCount == 1 && secretCount == 1 && deployCount < 1:
				namespaceDelete = append(namespaceDelete, namespace.GetName())
			default:
				//print deployments
				for _, deployment := range deployments.Items {
					fmt.Printf("\n🤖 Found namespace '%s' with deployments: %s \n", namespace.GetName(), deployment.GetName())
				}
			}

		} else {
			fmt.Printf("❗ Found a protected namespace: %s ⏩ \n", namespace.GetName())

		}
	}

	for _, namespace := range namespaceDelete {
		fmt.Printf("[To Delete] Namespace: %s\n", namespace)
	}

	out, err := cmd.Flags().GetBool("remove")
	if out {
		fmt.Println("Delete flag activated!")
	}

}
