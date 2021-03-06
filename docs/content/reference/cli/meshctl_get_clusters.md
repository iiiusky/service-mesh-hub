---
title: "meshctl get clusters"
weight: 5
---
## meshctl get clusters

Examine registered kubernetes clusters

### Synopsis

Examine registered kubernetes clusters

```
meshctl get clusters [flags]
```

### Options

```
  -h, --help   help for clusters
```

### Options inherited from parent commands

```
      --context string          Specify which context from the kubeconfig should be used; uses current context if none is specified
      --kube-timeout duration   Specify the default timeout for requests to kubernetes API servers (default 5s)
      --kubeconfig string       Specify the kubeconfig for the current command
  -n, --namespace string        Specify the namespace where Service Mesh Hub resources should be written (default "service-mesh-hub")
  -o, --output string           Output format. Valid values: [pretty, json, yaml] (default "pretty")
  -v, --verbose                 Enable verbose mode, which outputs additional execution details that may be helpful for debugging
```

### SEE ALSO

* [meshctl get](../meshctl_get)	 - Examine Service Mesh Hub resources

