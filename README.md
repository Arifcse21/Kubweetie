# Kubweetie
Kubernetes + Sweetie = Kubeweetie

<p>
Using KinD(Kubernetes in Docker) for local development.
KinD is a docker container runtime for Kubernetes.
</p>

### Kind setup
KinD Documentation[https://kind.sigs.k8s.io/docs/user/quick-start/]

```
kind create cluster --name <kind-cluster-name>
```

### Build socket server image
```
cd server
docker build -t socket_server:0.1.0 -f Dockerfile.python .
```

### Build socket client image
```
cd client
docker build -t socket_client:0.1.0 -f Dockerfile.go .
```

### Load built images to KinD(Kubernetes in Docker) - only for local development
```
kind load docker-image socket_server:0.1.0 --name <kind-cluster-name>
kind load docker-image socket_client:0.1.0 --name <kind-cluster-name>

# Check if images are loaded
docker exec -it <kind-container-name> crictl images
```

### Deploy socket server
```
kubectl apply -f server/server_deployment.yaml
```

### Deploy socket client
```
kubectl apply -f client/client_deployment.yaml
```

### Install Helm
```
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh
```

### Setup Kubernetes WebUI[https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/]
```
# Add kubernetes-dashboard repository
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/

# Deploy a Helm Release named "kubernetes-dashboard" using the kubernetes-dashboard chart
helm upgrade --install kubernetes-dashboard \
    kubernetes-dashboard/kubernetes-dashboard \
    --create-namespace --namespace kubernetes-dashboard
```

### Run Kubernetes WebUI
```
NAME: kubernetes-dashboard
NAMESPACE: kubernetes-dashboard
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
*************************************************************************************************
*** PLEASE BE PATIENT: Kubernetes Dashboard may need a few minutes to get up and become ready ***
*************************************************************************************************

Congratulations! You have just installed Kubernetes Dashboard in your cluster.

To access Dashboard run:
  `kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-kong-proxy 8443:443`

NOTE: In case port-forward command does not work, make sure that kong service name is correct.
      Check the services in Kubernetes Dashboard namespace using:
        kubectl -n kubernetes-dashboard get svc

Dashboard will be available at:
  https://localhost:8443

```

### Create Kubernetes Service Account
```
kubectl apply -f dashboard-admin-sa.yaml
```

### Role binding to Service Account
```
kubectl create clusterrolebinding dashboard-admin-binding \
    --clusterrole=cluster-admin \
    --serviceaccount=kubernetes-dashboard:dashboard-admin-sa
```

### Generate Kubernetes Auth Token
```
kubectl -n <NAMESPACE> create token <SERVICE_ACCOUNT>
```
