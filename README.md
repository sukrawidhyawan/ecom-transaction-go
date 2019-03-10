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
