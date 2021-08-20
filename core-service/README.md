[![Maintainability](https://api.codeclimate.com/v1/badges/afaeafb0caa35a6463f4/maintainability)](https://codeclimate.com/repos/611626fd92439c0161013db6/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/afaeafb0caa35a6463f4/test_coverage)](https://codeclimate.com/repos/611626fd92439c0161013db6/test_coverage)

# core-service
The core of jabar portal services

## Description


This project has 4 Domain layer :

* Domain Layer
* Repository Layer
* Usecase Layer
* Delivery Layer

### How To Run This Project

Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make test
```

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git@github.com:jabardigitalservice/portal-jabar-api.git

#move to project
$ cd core-service

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:7070

# Stop
$ make stop
```
