#!/bin/bash

set -e

IMAGE_NAME="http-test"

CGO_ENABLED=0 go build app.go
docker build -t $IMAGE_NAME .
