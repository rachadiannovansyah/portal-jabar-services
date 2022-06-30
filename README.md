[![Go Report Card](https://goreportcard.com/badge/github.com/jabardigitalservice/portal-jabar-services)](https://goreportcard.com/report/github.com/jabardigitalservice/portal-jabar-services)
[![Maintainability](https://api.codeclimate.com/v1/badges/e1b0eb219c1b35f76491/maintainability)](https://codeclimate.com/github/jabardigitalservice/portal-jabar-services/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e1b0eb219c1b35f76491/test_coverage)](https://codeclimate.com/github/jabardigitalservice/portal-jabar-services/test_coverage)
![ci workflow](https://github.com/jabardigitalservice/portal-jabar-services/actions/workflows/ci.yml/badge.svg)

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

# service-worker
The worker for provided core-service some cron-job process

## Description

Currently this service have 2 jobs :

* Archiving News
* Publishing News

### How To Run This Project

Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

#move to project
$ cd service-worker

# Build the docker image first
$ make docker

# Run the application
$ make run

# Stop
$ make stop
```