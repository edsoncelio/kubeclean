/*
Tool to clean empty namespaces on kubernetes!
*/

package test

import (
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func Test_getNamespaces(t *testing.T) {
	type args struct {
		clientset *kubernetes.Clientset
	}
	tests := []struct {
		name string
		args args
		want *v1.NamespaceList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNamespaces(tt.args.clientset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNamespaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
