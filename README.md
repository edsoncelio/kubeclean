# K8s namespace clean
Tool to remove empty namespaces on kubernetes

:warning: Do not run with a kubeconfig with role cluster-admin!

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edsoncelio/kubeclean)
![GitHub last commit](https://img.shields.io/github/last-commit/edsoncelio/kubeclean)

# Overview
![](example.png)

# Usage

## Requirements
* Go +1.15.2
* kubectl installed
* valid kubeconfig

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
 - [ ] add documentation
 - [x] check for deployment
 - [x] check for service
 - [x] check for statefulset
 - [ ] check for secret (beyond default)
 - [ ] check for service account (beyond default)
 - [ ] use external file to namespace exceptions (system namespaces)
 - [x] create the help flag
 - [ ] Configure CI (with github actions)
 
