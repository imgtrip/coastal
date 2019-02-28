#!/usr/bin/env bash

source ../scripts/platform.sh

ROOT_DIR=$(getRootDir)

cd ${ROOT_DIR}/build/docker/development

docker-compose up --build
