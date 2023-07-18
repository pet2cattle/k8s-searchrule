# k8s-searchrule

## Install


```
$ sudo cp ./kubectl-searchrule /usr/local/bin/
```

## Usage

```
$ kubectl searchrule --help
Usage of /usr/local/bin/kubectl-searchrule:
  -api-group string
    	API group to search for in role rules
  -g string
    	API group to search for in role rules (shorthand)
  -n string
    	Namespace to search for roles (shorthand)
  -namespace string
    	Namespace to search for roles
  -r string
    	Resource to search for in role rules (shorthand)
  -resource string
    	Resource to search for in role rules
  -v string
    	Verb to search for in role rules (shorthand) (default "*")
  -verb string
    	Verb to search for in role rules (default "*")
```


```
$ kubectl searchrule -n <namespace> -n <verb> -r <resource> -g <api-group>
```


$ kubectl searchrule -v '*' -r namespaces  -n pet2cattle-gitops
clusterrole/backplane-srep-admins-cluster
clusterrole/cluster-image-registry-operator
clusterrole/pet2cattle-dns-operator
clusterrole/pet2cattle-argocd-application-controller
clusterrole/pet2cattle-ingress-operator