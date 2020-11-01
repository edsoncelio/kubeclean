![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edsoncelio/kubeclean)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/edsoncelio/kubeclean)
![GitHub last commit](https://img.shields.io/github/last-commit/edsoncelio/kubeclean)
![Github workflow](https://github.com/edsoncelio/kubeclean/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/edsoncelio/kubeclean)](https://goreportcard.com/report/github.com/edsoncelio/kubeclean)

# kubeclean
Tool to remove empty namespaces on kubernetes

:warning: Do not run with a kubeconfig with role cluster-admin!

![](example.png)

## Requirements
* Go +1.15.2
* kubectl installed
* valid kubeconfig

## Installation 

### from release
Download the package from the [release page](https://github.com/edsoncelio/kubeclean/releases) and execute

### using go get   
`$ go get github.com/edsoncelio/kubeclean`

### from source
TODO

## Usage   
`$ kubeclean`

## TODO
 - [ ] add documentation
 - [ ] add tests
 - [x] check for deployment
 - [x] check for service
 - [x] check for statefulset
 - [ ] check for secret (beyond default)
 - [ ] check for service account (beyond default)
 - [ ] use external file to namespace exceptions (system namespaces)
 - [x] create the help flag
 - [x] configure CI (with github actions)
 - [x] configure release
 
