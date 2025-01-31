# fallback

Manually scale up your unavailable services

## Requirements

- access to kubeconfig (locally) or service account (in cluster)


## Usage

Deploy the following manifest

```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: fallback
  namespace: default

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: fallback-role
  namespace: default
rules:
  - apiGroups: ["apps"]
    resources: ["deployments"]
    verbs: ["get", "list", "update"]
  - apiGroups: [""]
    resources: ["services"]
    verbs: ["get"]
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["list"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: fallback-binding
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: fallback-role
subjects:
  - kind: ServiceAccount
    name: fallback
    namespace: default

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: fallback
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: fallback
  template:
    metadata:
      labels:
        app: fallback
    spec:
      serviceAccountName: fallback
      containers:
      - name: fallback
        image: ghcr.io/felipereyel/k8s-ingress-fallback:latest
        env:
        - name: USE_SA
          value: "true"
        - name: PORT
          value: "3000"

```

## Screens

### Home
![screenshot](screenshot/home.png)
