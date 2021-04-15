# go-starter-gin
---

## Description
符合Python开发者基于Django框架使用习惯 封装Go基于Gin框架 初始化工程
让开发者只关心 业务实现

- url 聚焦路由功能
- views 聚焦业务实现

## Feature

- 提供携程Apollo配置，提供统一配置管理
- 提供Gin Url分组管理，提供统一路由插拔策略
- 提供Mysql，提供统一数据库操作Client
- 提供Redis，提供统一Redis操作，Set Get
- 提供test 模块，熟悉如何完成业务开发

## 使用
```
apollo模块为全局所有配置唯一来源
修改apollo模块中 struct中相关配置(包含：mysql相关配置，redis相关配置，以及未来可能使用到的所有配置)

设置环境变量(正常Prod环境以及Dev Test 已经有统一的环境变量)，仅适用本地开发环境：

export RUNTIME_ENV=dev && export RUNTIME_CLUSTER=default && export RUNTIME_APP_NAME=op-robot-api && export LOG_BASE=debug

go run cmd/app/main.go 
```
## 验证使用
#### 请求 http://localhost:8080/v1/test 验证普通请求
![image](https://user-images.githubusercontent.com/81603118/113645745-0a8c9180-96ba-11eb-8c44-1e9e5fc4aa6c.png)

#### 请求 http://localhost:8080/v1/test1 验证apollo配置可行性
![image](https://user-images.githubusercontent.com/81603118/113645881-47588880-96ba-11eb-8786-e1cc1057943e.png)

#### 请求 http://localhost:8080/v1/test2 验证数据库查询接口可行性
![image](https://user-images.githubusercontent.com/81603118/113648966-24c96e00-96c0-11eb-89b8-3a93b0f31ca7.png)

#### 请求 http://localhost:8080/v1/test3 验证redis写入和查询是否正常
![image](https://user-images.githubusercontent.com/81603118/113668218-7e439400-96e4-11eb-8f82-b8981e6d32f7.png)

### 请求 http://localhost:8080/swagger/index.html 验证swagger是否正常
![image](https://user-images.githubusercontent.com/81603118/114156869-9b32ce00-9955-11eb-821c-3e184e077886.png)

### 验证build
![image](https://user-images.githubusercontent.com/81603118/114799570-7792dc00-9dca-11eb-943a-9b74c1b531d2.png)
![image](https://user-images.githubusercontent.com/81603118/114799627-a6a94d80-9dca-11eb-90a4-f531a64331a9.png)





