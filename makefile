#!/bin/bash

export LANG=zh_CN.UTF-8

ENV_ARG=
LINUX_ARG=GOOS=linux GOARCH=amd64
BUILD_ARG=-ldflags "-s -X github.com/beanscc/rango/utils.buildTime=`date '+%Y-%m-%dT%H:%M:%S%z'` -X github.com/beanscc/rango/utils.gitHash=`git rev-parse HEAD`"

PRO_ROOT=$(CURDIR) # project root
BIN_DIR=${PRO_ROOT}/bin # bin dir
#CMD_DIR=$(PRO_ROOT)/cmd # cmd dir

pwd:
	echo ${BIN_DIR}

build:
	cd cmd/autoscheme; $(ENV_ARG) go build $(BUILD_ARG) -o ../../bin/autoscheme

  #	build linux version
	cd cmd/autoscheme; $(ENV_ARG) $(LINUX_ARG) go build $(BUILD_ARG) -o ../../bin/autoscheme_linux

clean:
	rm bin/autoscheme
  # rm linux
	rm bin/autoscheme_linux