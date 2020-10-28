# K8s namespace clean
Tool to remove empty namespaces on kubernetes

:warning: Do not run with a kubeconfig with role cluster-admin or set the protected namespaces!

![GitHub](https://img.shields.io/github/license/edsoncelio/kubeclean)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edsoncelio/kubeclean)

# Usage

## Requirements
* Go +1.15.2

## Installation
Inside the app directory, build:  
`go build -o ./kubeclean`

And run (getting kubeconfig from default path):   
`./kubeclean`

or, passing the absolute path to kubeconfig file:   
`./kubeclean --kubeconfig /my/kubeconfig/file`

To get help:   
`./kubeclean --help`


# Documentation
TODO

# TODO
 -  [ ] add documentation
 - [x] check for deployment
 - [x] check for service
 - [x] check for statefulset
 - [ ] check for secret (beyond default)
 - [ ] check for service account (beyond default)
 - [ ] use external file to namespace exceptions (system namespaces)
 - [x] create the help flag
 
