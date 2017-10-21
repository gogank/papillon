#!/usr/bin/env bash


PROJECT_PATH="${GOPATH}/src/github.com/gogank/papillon"

# output root dir
DUMP_PATH="${PROJECT_PATH}/build"

# config file path
CONF_PATH="${PROJECT_PATH}/configuration/blog"

rm -rf ${DUMP_PATH}
cp -r ${CONF_PATH} ${DUMP_PATH}/
cd ${PROJECT_PATH} && govendor build -ldflags -s -o ${DUMP_PATH}/papi -tags=embed