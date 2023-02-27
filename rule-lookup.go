package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var namespace string
	var verb string
	var resource string
	var apiGroup string

	// Define flags
	flag.StringVar(&namespace, "namespace", "", "Namespace to search for roles")
	flag.StringVar(&namespace, "n", "", "Namespace to search for roles (shorthand)")
	flag.StringVar(&verb, "verb", "*", "Verb to search for in role rules")
	flag.StringVar(&verb, "v", "*", "Verb to search for in role rules (shorthand)")
	flag.StringVar(&resource, "resource", "", "Resource to search for in role rules")
	flag.StringVar(&resource, "r", "", "Resource to search for in role rules (shorthand)")
	flag.StringVar(&apiGroup, "api-group", "", "API group to search for in role rules")
	flag.StringVar(&apiGroup, "g", "", "API group to search for in role rules (shorthand)")

	// Parse flags
	flag.Parse()

	if resource == "" {
		fmt.Println("resource cannot be empty")
		os.Exit(1)
	}

	homedir, _ := os.UserHomeDir()

	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: homedir + "/.kube/config"}, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		fmt.Printf("ERROR: can't create config from user's kubeconfig: %s\n", err.Error())
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

	if namespace != "" {
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
