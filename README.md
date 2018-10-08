# Commenting API Service on GAE

An commenting API service which is run on GAE, based on Clean Architecture and Domain Driven Design

## About this project

This project is built in order to learn about followings.
- Clean Architecture
- Domain Driven Design
- MicroServices Architecture (just a little)

This project keeps these principles as much as possible.  
So this project has some over-engineered parts.

## Application Architecture Overview

![](doc/architecture-overview.png)

- This service has 3 bounded contexts.
  - Commenting context
  - Auth context
  - Notification context
- Commenting API is based on Clean Architecture.

### TODO

- Separate Commenting API Service and Notification Service into different GAE service

## Domain Driven Design in this project

### Bounded contexts

#### Commenting context (core domain)

Post a comment, delete a comment and view posted comments.

#### Notification context (generic sub domain)

Notify about some events to users or administrators.  
In this project Google Cloud Pub/Sub are used for publishing/subscribing domain event.
This domain subscribes domain events and then notifies to appropriate persons.

#### Auth context (generic sub domain)

This domain helps user authentication.  
In this project Firebase Authentication Service has auth domain role.

## Clean Architecture in Commenting API

### Architecture overview

To be written

### Benefits of Clean Architecture

Because of Clean Architecture, application logic and domain logic are independent with detail of infrastructure.  
Followings are not appeared in the core of the application.

- Various packages which is related to Google App Engine infrastructure
- Technological details of web application (e.g. context.Context)

# Setup

## Requirements

- go 1.8
- google-cloud-sdk
    - goapp
    - dev_appserver.py
- dep

## Setup for local development

- Rewrite yaml for your environment
- Resolve dependencies
    ```shell
    # /path/to/comment-api-on-gae/src/commenting
    $ GOPATH=/path/to/comment-api-on-gae dep ensure
    ```

## Run

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae dev_appserver.py app --enable_watching_go_path --log_level=debug --datastore_path=.storage
```


## Test

todo

## Deploy

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae goapp deploy app
```

## Setup Cloud Services

### Google App Engine

- Create datastore index
- Issue service account which have following privileges
   - Firebase data manager
   - Pubsub editor
- Add authorized mail sender on Google App Engine console

### Firebase

- Enable anonymous login

### Google cloud pubsub

- Create topic `domain-event`
    ```
    gcloud beta pubsub topics create domain-event
    gcloud beta pubsub subscriptions create comment-api-domain-event \
        --topic domain-event \
        --push-endpoint \
        https://YOUR_PROJECT_ID.appspot.com/_ah/push-handlers/domain-event \
        --ack-deadline 10
    ```
