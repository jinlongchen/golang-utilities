#!/usr/bin/env bash

CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-s -w -X github.com/jinlongchen/golang-utilities/config/aesKeyKey=J^_^inlongChen" && ./aes && rm ./aes

