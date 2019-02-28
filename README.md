# IMGTRIP

[![pipeline status](https://gitlab.com/qskane/imgtrip/badges/master/pipeline.svg)](https://gitlab.com/qskane/imgtrip/commits/master)

## DOWNLOAD

**将需要子模块访问权限,子模块见`.gitmodules` 文件**

1. 下载本仓库  `git clone git@gitlab.com:qskane/imgtrip.git`
2. 下载子模块  `git submodule update`

## INSTALL

### require

> 以下为推荐版本，不作强制限制，若需要，可自行尝试其他版本

- docker >= 17.0
- go >= 1.10

### linux平台

1. 启动数据库

    ```bash
    cd ./build/package

    # 启动docker container,包含: imgtrip-db, imgtrip-test-db, adminer
    docker-compose up -d
    ```

    浏览器访问 **http//localhost:8306**,可使用对应帐号密码登陆即可
    > database: imgtrip-db || imgtrip-test-db
    > user:root
    > password:root

2. 启动前

    ```bash

    # 可选,安装依赖
    dep ensure

    cd ./cmd/imgtrip_migrate
    go run main.go

<<<<<<< HEAD
    ```
=======
See the `/pkg` directory for examples.

### `/vendor`
>>>>>>> e82886e... more references

//......................
// 未完成
//......................


3. 启动

    ```bash
    cd ./user

    # 安装golang依赖
    # 依赖管理目前未找到解决方案,在下一步中提示缺少库时自行手动安装: go get github.com/YOUR-MISSING/REPOSITORY

<<<<<<< HEAD
    # 启动user-service
    # 将包含： 50051作为grpc内部通信服务端口 , 9090作为http访问端口(其他端口见./hub/conf/conf.go)
    go run main.go
    ```
=======
See the `/api` directory for examples.

## Web Application Directories

### `/web`

Web application specific components: static web assets, server side templates and SPAs.

## Common Application Directories

### `/configs`

Configuration file templates or default configs.

Put your `confd` or `consule-template` template files here.

### `/init`

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

### `/scripts`

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple (e.g., `https://github.com/hashicorp/terraform/blob/master/Makefile`).

See the `/scripts` directory for examples.

### `/build`

Packaging and Continous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.

Put your CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory.

### `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).

### `/test`

Additional external test apps and test data. Feel free to structure the `/test` directory anyway you want. For bigger projects it makes sense to have a data subdirectory (e.g., `/test/data`).

See the `/test` directory for examples.

## Other Directories

### `/docs`

Design and user documents (in addition to your godoc generated documentation).

See the `/docs` directory for examples.

### `/tools`

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.

See the `/tools` directory for examples.

### `/examples`

Examples for your applications and/or public libraries.

See the `/examples` directory for examples.

### `/third_party`

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

Other assets to go along with your repository.

## Directories You Shouldn't Have

### `/src`

Some Go projects do have a `src` folder, but it usually happens when the devs came from the Java world where it's a common pattern. If you can help yourself try not to adopt this Java pattern. You really don't want your Go code or Go projects to look like Java :-)


## Badges

* [Go Report Card](https://goreportcard.com/) - It will scan your code with `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Replace `github.com/golang-standards/project-layout` with your project reference.

* [GoDoc](http://godoc.org) - It will provide online version of your GoDoc generated documentation. Change the link to point to your project.

* Release - It will show the latest release number for your project. Change the github link to point to your project.

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notes

A more opinionated project template with sample/reusable configs, scripts and code is a WIP.
>>>>>>> e82886e... more references

    启动界面后无报错，进程挂起，提示正在监听对应端口，表示成功完成
