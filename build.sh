#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build main.go

#docker build . -t main:latest
