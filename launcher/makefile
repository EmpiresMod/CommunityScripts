#!/usr/bin/make -f

SHELL=/bin/sh
bin=bin
name=launcher

all: build

build:
	go build -v -o bin/$(name)

windows:
	GOOS=windows go build -v -ldflags="-H windowsgui" -o="bin/$(name).exe"

clean:
	go clean -x

remove:
	go clean -i
