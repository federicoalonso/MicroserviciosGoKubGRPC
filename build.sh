#!/bin/bash

go clean --cache && go test -v -cover ../MicroserviciosGoKubGRPC/authentication/...
go build -o authentication/authsvc authentication/main.go