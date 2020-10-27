# Kubernetes namespace clean
Tool to remove empty namespaces on kubernetes

:warning: Do not run with a kubeconfig with role cluster-admin or set the protected namespaces!

![](https://img.shields.io/github/license/edsoncelio/kubeclean)

# Usage

## Requirements
An existing kubeconfig file on `./kube/config`

## Installation
Inside the app directory, build:  
`go build -o ./kubeclean`

And run:   
`./kubeclean`


# Documentation
TODO

# TODO
 -  [ ] add documentation
 - [x] check for deployment
 - [ ] check for service
 - [x] check for statefulset
 - [ ] check for secret (beyond default)
 - [ ] check for service account (beyond default)
 - [ ] use external file to namespace exceptions
 - [x] create the help flag
 
