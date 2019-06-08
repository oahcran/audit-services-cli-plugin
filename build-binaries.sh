#!/bin/bash

PLUGIN_NAME=audit-services
PLUGIN_VERSION=0.1.0

GOOS=linux GOARCH=amd64 go build -o ${PLUGIN_NAME}_${PLUGIN_VERSION}.linux64
#GOOS=linux GOARCH=386 go build -o ${PLUGIN_NAME}_${PLUGIN_VERSION}.linux32
GOOS=windows GOARCH=amd64 go build -o ${PLUGIN_NAME}_${PLUGIN_VERSION}.win64
#GOOS=windows GOARCH=386 go build -o ${PLUGIN_NAME}_${PLUGIN_VERSION}.win32
GOOS=darwin GOARCH=amd64 go build -o ${PLUGIN_NAME}_${PLUGIN_VERSION}.osx