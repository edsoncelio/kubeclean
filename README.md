# kubeclean
Tool to remove empty namespaces on kubernetes

## How to use
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
 
