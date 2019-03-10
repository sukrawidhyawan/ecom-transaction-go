# ecom-transaction-go

This is repository for transaction microservice of Ecom project build with Go and uses Postgres and GRPC

## Running service locally

- go to the specific services' directory

- building the image:
  `docker build -t ecom-transaction:1.0 .` (change the name of the build name as convenience)

- running the image
  `docker run --rm -p 8000:8000 ecom-transaction:1.0` (change the name of the build name as convenience)
  or
  `docker run --rm -t -i -p 8000:8000 ecom-transaction:1.0`

- tagging the image
  `docker tag ecom-transaction:1.0 alhiee/ecom-transaction:1.0`

- pushing the image to docker hub
  `docker push alhiee/ecom-transaction:1.0`

## GCP Kubernetes

tutorial: https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app

- setting up kubernetes
  `gcloud container clusters get-credentials {{standard-cluster-1}} --region {{asia-southeast1-a}}`

- running kubernetes for the service
  `kubectl run {{ecom-transaction}} --image={{registry.hub.docker.com/alhiee/ecom-transaction:1.0}} --port {{8000}}`

- exposing the service to internet
  `kubectl expose deployment {{ecom-transaction}} --type=LoadBalancer --port {{80}} --target-port {{8000}}`

- getting the external IP to access the kubernetes and other details
  `kubectl get service`
