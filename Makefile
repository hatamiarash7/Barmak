SHELL := /bin/bash

## help: show this help message
help:
	@ echo -e "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## build: build app's binary
build:
	@ go build -a -installsuffix cgo -o Barmak .

## run: run the app
run: build
	@ ./Barmak