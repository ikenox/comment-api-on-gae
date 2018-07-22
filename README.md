# Comment API on GAE

API server of commenting service which is run on GAE

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
