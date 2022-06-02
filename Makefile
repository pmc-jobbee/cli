# SHELL := $(shell echo $$SHELL)

VERSION=$(shell cat VERSION)

all: build test

build:
	make -C cli

clean:
	make -C cli clean

test:
	go test -v ./...

tidy:
	go mod tidy