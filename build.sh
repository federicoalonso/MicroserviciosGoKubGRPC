#!/bin/bash

go clean --cache && go test -v -cover ../MicroserviciosGoKubGRPC/...
go build -o authentication/authsvc authentication/main.go