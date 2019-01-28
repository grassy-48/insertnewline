#! /usr/bin/make
ifeq ($(OS),Windows_NT)
	BUILD_TARGET_FILES = insernewline.exe main.go
else
	BUILD_TARGET_FILES ?= insertnewline main.go
endif

.DEFAULT_GOAL := prepare

all: prepare build

prepare: depend-test depend 

depend:
	@vgo build

depend-test:
	@vgo test all

build:
	@go build -o $(BUILD_TARGET_FILES)

build-closs:
	@GOOS=windows GOARCH=386 vgo build -o insernewline_win.exe main.go
	@GOOS=darwin GOARCH=amd64 vgo build -o insernewline_mac main.go

run:
	@vgo run main.go
