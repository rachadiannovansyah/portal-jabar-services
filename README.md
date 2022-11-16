[![Go Report Card](https://goreportcard.com/badge/github.com/jabardigitalservice/portal-jabar-services)](https://goreportcard.com/report/github.com/jabardigitalservice/portal-jabar-services)
[![Maintainability](https://api.codeclimate.com/v1/badges/e1b0eb219c1b35f76491/maintainability)](https://codeclimate.com/github/jabardigitalservice/portal-jabar-services/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e1b0eb219c1b35f76491/test_coverage)](https://codeclimate.com/github/jabardigitalservice/portal-jabar-services/test_coverage)
<br>
![ci workflow](https://github.com/jabardigitalservice/portal-jabar-services/actions/workflows/ci.yml/badge.svg)
[![GitHub issues](https://img.shields.io/github/issues/jabardigitalservice/portal-jabar-services)](https://github.com/jabardigitalservice/portal-jabar-services/issues)
<br>
[![GitHub stars](https://img.shields.io/github/stars/jabardigitalservice/portal-jabar-services)](https://github.com/jabardigitalservice/portal-jabar-services/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/jabardigitalservice/portal-jabar-services)](https://github.com/jabardigitalservice/portal-jabar-services/network)

# core-service
The core of jabar portal services

## Tech Stacks

- **Golang** - <https://go.dev//>
- **Echo Framework** - <https://echo.labstack.com//>
- **Elasticsearch** - <https://www.elastic.co//>
- **MySQL** - <https://www.mysql.com/>
- **Redis** - <https://redis.io/>

## Description

This project has 4 Domain layer :

* Domain Layer
* Repository Layer
* Usecase Layer
* Delivery Layer

## Repo Structure core-service
```
├── .github/            * all workflows github actions
  └── workflows/
├── docker/
├── src/
  └── cmd/              * all the command of app here.
  └── config/           * contains config like db, aws, redis, etc.
  └── database/         * database migrations etc.
  └── domain/           * all the contract, struct, interface etc.
  └── helpers/          * contains a helpers function etc.
  └── utils/            * contains utilities of email, conn, apm etc.
  └── policies/         * policies of authorized user.
  └── middleware/       * request's middleware.
  └── modules/          * contains modules of application.
    └── <delivery>
      └── <http>
          └── <name>_handler.go   * contains file of delivery handler.
    └── <repository>
      └── <mysql>
          └── <name>_mysql.go     * contains file of repository mysql, etc.
    └── <usecase>
      └── <name>_ucase.go         * contains file of busines logic / usecase.
└── ...
```

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

## Repo Structure service-worker
```
├── docker/
├── src/
  └── cmd/             * all the command of app here.
  └── config/          * contains config like db, aws, redis, etc.
  └── job/             * job functions.
  └── utils/           * contains utilities of email, conn, apm etc.
```

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
