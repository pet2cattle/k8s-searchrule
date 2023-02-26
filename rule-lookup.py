import os
import argparse
from kubernetes import client, config

def main():
    # Define flags
    parser = argparse.ArgumentParser(description='rule searcher')
    parser.add_argument('--namespace', default='', help='Namespace to search for roles')
    parser.add_argument('--verb', default='*', help='Verb to search for')
    parser.add_argument('--resource', required=True, help='Resource to search for')
    parser.add_argument('--api-group', default='', help='API group to search for')
    args = parser.parse_args()

    if args.resource == "":
        print("resource cannot be empty")
        exit(1)

    config.load_kube_config()

    # List all ClusterRoles and Roles in the given namespace
    clusterRoleList = client.RbacAuthorizationV1Api().list_cluster_role()
    roleList = []
    if args.namespace != '':
        roleList = client.RbacAuthorizationV1Api().list_namespaced_role(args.namespace)

    # Check each ClusterRole for matching rules
    for role in clusterRoleList.items:
        if rulematching(role.rules, args.api_group, args.resource, args.verb):
            print(f"clusterrole/{role.metadata.name}")

    # Check each Role for matching rules
    for role in roleList.items:
        if rulematching(role.rules, args.api_group, args.resource, args.verb):
            print(f"role/{role.metadata.name}")

def rulematching(rules, apiGroup, resource, verb):
    if not rules:
        return False
    for rule in rules:
        if not rule.api_groups:
            continue
        if apiGroup and apiGroup not in rule.api_groups:
            continue
        if not rule.resources:
            continue
        if resource not in rule.resources and '*' not in rule.resources:
            continue
        if verb not in rule.verbs and '*' not in rule.verbs:
            continue
        return True
    return False

if __name__ == '__main__':
    main()