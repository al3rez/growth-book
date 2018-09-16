#!/usr/bin/env bash

go get -v ./...
go get -v github.com/amacneil/dbmate

dbmate wait && dbmate up
go run main.go
