# Tell Makefile to use bash
SHELL := /bin/bash

.PHONY: init up up-dev down install install-dev cache clear key seed bash

generate-swagger:
	sudo swagger generate spec -o ./swagger.json;

run-swagger:
	docker pull swaggerapi/swagger-editor;
	docker run --rm -d -p 80:8080 swaggerapi/swagger-editor;

migration:
	go run migration/main.go;

test:
	go test -v ./services/...

start:
	gin -p 8001 -a 3000 main.go;