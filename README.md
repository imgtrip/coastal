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


3. 启动

    ```bash
    cd ./user

    # 安装golang依赖
    # 依赖管理目前未找到解决方案,在下一步中提示缺少库时自行手动安装: go get github.com/YOUR-MISSING/REPOSITORY

    # 启动user-service
    # 将包含： 50051作为grpc内部通信服务端口 , 9090作为http访问端口(其他端口见./hub/conf/conf.go)
    go run main.go
    ```

    启动界面后无报错，进程挂起，提示正在监听对应端口，表示成功完成
