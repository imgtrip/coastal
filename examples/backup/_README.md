
![avatar](./ts/static/logo/circular-160.png) 

## INSTALL

### require

> 以下为推荐版本，不作强制限制，若需要，可自行尝试其他版本

- docker >= 17.0
- go >= 1.10
- typescript >= 2.8

### linux平台

1. 启动数据库
    
    ```bash
    # 切换到根目录
    cd /path/to/repository
    
    # 启动docker,包含 user-db,img-db,log-db,adminer 四个container
    docker-compose up -d
    ```
    浏览器访问**http//localhost:8306**,可使用对应帐号密码登陆表示成功完成
    > database: user-db || img-db || log-db  
    > user:root  
    > password:root    

2. 启动services

    **启动user-service**
    ```bash
    cd ./user
    
    # 安装golang依赖
    # 依赖管理目前未找到解决方案,在下一步中提示缺少库时自行手动安装: go get github.com/YOUR-MISSING/REPOSITORY
 
    # 启动user-service
    # 将包含： 50051作为grpc内部通信服务端口 , 9090作为http访问端口(其他端口见./hub/conf/conf.go)
    go run main.go
    ```
    
    若需要启动其他service 参考**user-service**启动方法
    其他service: 
    - **img-service** 
    - **log-service**
    
    启动界面后无报错，进程挂起，提示正在监听对应端口，表示成功完成

3. 启动前端项目

    本例使用yarn,可自行选择npm或其他工具
    
    ```bash
    cd ./ts
    
    # 安装依赖
    sudo yarn init
    sudo yarn
    
    # 开发模式启动
    sudo yarn dev
    ```
    浏览器访问 **http://localhost:3000**,能显示任何非报错界面即表示完成


## 开发

目录树
- hub 各service复用工具项目
- user user相关服务项目
- img  img相关服务项目
- log  log相关服务项目
- ts   前端项目
- proto rpc交互存根

查看`./FEATURES.md`获取开发进度

### 后端开发

1. 书写proto, 编辑`./proto`中对应文件

2. 编译proto

    ```bash
    sh ./complire.sh
    ```
    将重新生成`typescript`,`golang`所需要的编译后文件


3. 编写对应后端功能,实现逻辑

4. 重启service

### 前端开发

1. 书写纯视图逻辑
2. 参见`./proto`中对应文件存根，获取对应数据交互API
