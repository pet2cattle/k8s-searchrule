package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("Usage: %s <namespace> <verb> <resource> <api-group>\n", os.Args[0])
	}

	namespace := os.Args[1]
	verb := os.Args[2]
	resource := os.Args[3]
	apiGroup := os.Args[4]

	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	log.Fatalf("Failed to create Kubernetes config: %v\n", err)
	// }

	homedir, _ := os.UserHomeDir()

	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: homedir + "/.kube/config"}, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		fmt.Printf("%s: can't create config from kubeConfigPath: %s\n", time.Now().Format(time.RFC3339), err.Error())
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v\n", err)
	}

	ctx := context.TODO()
	opts := metav1.ListOptions{}

	// List all ClusterRoles and Roles in the given namespace
	clusterRoleList, err := clientset.RbacV1().ClusterRoles().List(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to list ClusterRoles: %v\n", err)
	}

	// Check each ClusterRole for matching rules
	for _, role := range clusterRoleList.Items {
		if rulematching(role.Rules, apiGroup, resource, verb) {
			fmt.Printf("clusterrole/%s\n", role.Name)
		}
	}

	roleOpts := metav1.ListOptions{}
	roleList, err := clientset.RbacV1().Roles(namespace).List(ctx, roleOpts)
	if err != nil {
		log.Fatalf("Failed to list Roles: %v\n", err)
	}

	// Check each Role for matching rules
	for _, role := range roleList.Items {
		if rulematching(role.Rules, apiGroup, resource, verb) {
			fmt.Printf("role/%s\n", role.Name)
		}
	}
}

func rulematching(rules []v1.PolicyRule, apiGroup string, resource string, verb string) bool {
	for _, rule := range rules {
		for _, apiGroupEntry := range rule.APIGroups {
			if apiGroupEntry == apiGroup {
				for _, resourceEntry := range rule.Resources {
					if resourceEntry == resource || resourceEntry == "*" {
						for _, verbEntry := range rule.Verbs {
							if verbEntry == verb || verbEntry == "*" {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}
