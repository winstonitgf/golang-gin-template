# Tell Makefile to use bash
SHELL := /bin/bash

.PHONY: init up up-dev down install install-dev cache clear key seed bash

swag:
	swag init
	
test:
	go test -v ./service/...

start:
	gin -p 8001 -a 3000 main.go;