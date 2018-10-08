# Commenting Service on GAE

An commenting API service using Google App Engine and some Google Cloud Platform Services

## About this project

This project aims to learn about followings.
- Clean Architecture
- Domain Driven Design
- MicroServices Architecture (just a little)

This project keeps these principles as much as possible.  
So there are several parts which is over-engineered.

## Application Architecture Overview

## Domain Driven Design

### Bounded contexts in Commenting Service

#### Commenting domain (core domain)

Post a comment, delete a comment and view posted comments.

#### Notification domain (generic sub domain)

Notify about some events to users or administrators.  
In this project Google Cloud Pub/Sub are used for publishing/subscribing domain event.
This domain subscribes domain events and then notifies to appropriate persons.

#### Auth domain (generic sub domain)

This domain helps user authentication.  
In this project Firebase Authentication Service has auth domain role.

## A benefits of Clean Architecture

Because of Clean Architecture, application logic and domain logic are independent with detail of infrastructure.  
Followings are not appeared in the core of the application.

- Various packages which is related to Google App Engine
- Technological details of web application (e.g. context.Context)

## Requirements

- go 1.8
- google-cloud-sdk
    - goapp
    - dev_appserver.py
- dep

## Setup

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae dep ensure
```

## Run

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae dev_appserver.py app --enable_watching_go_path --log_level=debug --datastore_path=.storage
```

## Deploy

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae goapp deploy app
```

## Setup cloud services

### Google App Engine

- Create datastore index
- issue service account which have following privileges
   - firebase data manager
   - pubsub publisher/subscriber
- add authorized mail sender on appengine console
- rewrite yaml

### Firebase

- enable anonymous login

### Google cloud pubsub

- create topic

gcloud beta pubsub topics create domain-event
gcloud beta pubsub subscriptions create comment-api-domain-event \
    --topic domain-event \
    --push-endpoint \
    https://YOUR_PROJECT_ID.appspot.com/_ah/push-handlers/domain-event \
    --ack-deadline 10
