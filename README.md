# k8s-searchrule

## Install

Copy the binary to any folder within the `PATH`:

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

Search for namespaced resources:

```
$ kubectl searchrule -n <namespace> -n <verb> -r <resource> -g <api-group>
```

Search for cluster resouces:

```
$ kubectl searchrule -n <verb> -r <resource> -g <api-group>
```