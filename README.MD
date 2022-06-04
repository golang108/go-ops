# Go-OPS

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Go-ops是面向云平台系统, 基于 PaaS 技术支持客户实现 DevOps 解决方案的平台。提供完善作业管理、配置管理、定期调度、健康巡检等模块，助力运维人员快速、低成本、自动化的管理配置运维系统。



## 功能
 - 节点管理
   - 节点连接/断开
   - 节点健康状态
   - 节点自我升级
 - 资源管理
    - 脚本资源
    - 巡检项
    - 巡检模板
    - 文件资源
    - 插件资源
 - 任务管理
   - 脚本任务
   - 文件分发
   - 巡检任务
   - 预设任务
   - 定时任务
 - agent管理
 - 作业管理
    - 分批作业
    - 作业编排
 - APP(对接方)管理
 - 权限控制
   - 用户权限
   - 对接方权限
 - 远程连接
   - ssh协议
   - vnc协议


## 构建项目

- make
  
  > make build



- 配置文件说明
  - config.json  agent配置文件
  ```json
  {
    "port":13333,  // agent 监听的端口号
    "bootlist":["tcp://127.0.0.1:9999"]  // 引导网络地址
  }
  ```
  - config.yaml  apiserver配置文件
  
  ```yaml

  database:
    default:
      link:  "mysql:root:root@tcp(127.0.0.1:3306)/ops" # 数据库连接串
      debug: "true"                                    # 是否开启调试模式

  server:
    ServerRoot: "public"        # 静态资源文件目录
    address: ":8199"            # api 监听端口
    openapiPath: "/api.json"    # openapi 接口文件
  
  peer:
    port: 9999                # 节点监听端口号
    boots:                    # 加入的网络节点列表
      - addr: "tcp://xxxx"    # 引导节点地址1
      - addr: "tcp://xx11122" # 引导节点地址2
  ```


## 运行项目

> 如果是第一次运行需要在数据库中创建数据库，并导入scripts/ops.sql  这个文件的表结构信息

- apiserver
  > bin/apiserver

  或者
  >  go run cmd/server/main.go

- agent
  > bin/agent
  
  或者
  > agent: go run cmd/agent/main.go


## 前端

https://github.com/i4de/opsApp


## 文档

https://www.yuque.com/lewisxw/goops