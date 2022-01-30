The repository contains the solution for the [challenge](./Challenge.md) in Go and Python.

# Steps to run
**Perquisite: Kubernetes**

## Python API
This service will be available on port `3035`
```sh 
cd py && sh run.sh
``` 

## Go API
This service will be available on port `3030`
```sh 
cd go && sh run.sh
``` 

## What does run.sh do?
This file will create the namespace, deployment and service for the respective application on K8s.

The application will be served under namespace `web-tech-challenge-saurabhyeolekar`