# Deploy Golang App to GKE

## Install and Configure GCloud CLI

- https://cloud.google.com/sdk/docs/install

## Configure Google Container Registry

- https://cloud.google.com/sdk/gcloud/reference/auth/configure-docker
- https://cloud.google.com/container-registry/docs/pushing-and-pulling?hl=pt_br

## Build Image Tag

```bash
docker tag ${PROJECT_NAME} gcr.io/${GCLOUD_PROJECT_ID}/${PROJECT_NAME}
```

## Push Image to GCP

```bash
docker push gcr.io/${GCLOUD_PROJECT_ID}/${PROJECT_NAME}
```

## Create Cluster

- Connect to Cluster (Get command on GCP Console)

## Apply Deployment

```bash
kubctl apply -f k8s/
````

## Get Ingress IP

```bash
kubectl get ingress ${PROJECT_NAME}-ingress --output yaml
```

- output:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
/* omit data... */
spec:
  rules:
  - http:
      paths:
      - backend:
          service:
            name: go-sandbox-service
            port:
              number: 8089
        path: /
        pathType: Prefix
status:
  loadBalancer:
    ingress:
    - ip: 34.111.39.23
```

