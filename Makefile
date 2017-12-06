SHELL := /bin/bash

.PHONY: test
test: 
	go test -v -race ./...

