# Kubernetes Empty Namespace Clean
Tool to remove empty namespaces on kubernetes

![](https://img.shields.io/github/license/edsoncelio/kubeclean)

# Usage

## Requirements
An existing kubeconfig file on `./kube/config`

## Installation
Inside the app directory, build:  
`go build -o ./kubeclean`

And run:   
`./kubeclean`

# TODO
 -  [ ] add documentation
 - [x] check for deployment
 - [ ] check for service
 - [ ] check for statefulset
 - [ ] check for secret (beyond default)
 - [ ] check for service account (beyond default)
 - [ ] use external file to namespace exceptions
 - [ ] create the help flag
 
