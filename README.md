# Deployer [![GoDoc](https://godoc.org/https://github.com/izner32/deployer?status.svg)](https://godoc.org/https://github.com/izner32/deployer) [![Go Report Card](https://goreportcard.com/badge/https://github.com/izner32/deployer)](https://goreportcard.com/report/https://github.com/izner32/deployer)

## File Structure
```
    .github: contains the ci workflow 
    cmd: code implementation for commands
        │
        │―――― deploy.go: code for respective command
        │―――― doc.go: contains package declaration details
        │―――― root.go: code for the root command (in our case, deployer)
        │―――― status.go: code for respective command
        │―――― undeploy.go: code for respective command
        │―――― version.go: code for respective command
    .gitignore
    LICENSE
    main.go: entrypoint to the program
    README.md
    go.mod: details about the modules used; auto-generate ./go.sum
```

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN="$(pwd)" go install https://github.com/izner32/deployer@latest
```

## Command Structure
deployer {command} {sub-command} {flag}

Example
```
deployer deploy web 
```

## Deployer API
#### Deployment Resources
* deploy: Deploy the application
    * web:
    * api: 
    * database:
* undeploy: Remove deployment of the application
    * web: 
    * api: 
    * database: 

#### Other Resources
* status: Get status of the current deployment
    * web: 
    * api: 
    * database: 
* version: Get Version of the Deploy CLI

#### Flags
Global Flags
* config: 

## Screenshots

```
```
