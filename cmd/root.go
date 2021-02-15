package cmd

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "kubeclean",
		Short: "A CLI to remove empty kubernetes namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
			execNamespaceCheck(kubeconfig)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	home, _ := homedir.Dir()
	if home != "" {
		rootCmd.PersistentFlags().StringP("kubeconfig", "k", filepath.Join(home, ".kube", "config"), "kubeconfig file (default is $HOME/.kube)")
	}
}
