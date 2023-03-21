# GCloud 开发手册

## 一、开发环境

操作系统：Windows 11

开发工具：VSCode、Docker desktop(Remote、MySQL、Redis)

前端：Node 14+、yarn 1.18.22、Vite 2.4.2

后端：Go 1.18+ 

## 二、前端

### 技术栈

- 框架：Vue3.2、Pinia
- UI库：Naive UI、tailwindcss、vue3-lottie、animate.css、vicons、sass
- 开发：Vite、Typescript
- 以及各种工具

### 目录结构

聪明的你一看就懂！



## 三、后端

### 调试命令

```shell
$ cd core
# 启动Api服务
$ go run core.go -f etc/core-api.yaml

# 使用api文件生成代码
$ goctl api go -api core.api -dir . -style go_zero
```

### 配置

#### 系统环境变量

> 自行Google：Windows配置环境变量、Linux配置环境变量

| 变量名           | 类型   | 备注                      |
| ---------------- | ------ | ------------------------- |
| MailPassword     | string | 邮箱授权码（注册服务）    |
| RedisPassword    | string | Redis 密码（验证码）      |
| TencentSecretKey | string | COS SecretKey（文件上传） |
| TencentSecretID  | string | COS SecretID（文件上传）  |

如何使用环境变量？代码详见：[define.go](/core/define/define.go)



#### 邮箱注册配置

why：提供邮箱发送验证码注册功能

目标：开启邮箱 SMTP 服务并获取**密钥**

示例：网易邮箱 (@163.com)

步骤：

1.注册网易邮箱账号：https://email.163.com/；

2.进入控制台点击顶部导航栏中的“**设置**”，弹出下拉菜单中点击“**POP3/SMTP/IMAP**”项；

3.开启服务 “**IMAP/SMTP服务**”，并将获取到的授权码保存在主机环境变量中即可。



#### 对象存储 COS 配置

why：存储用户文件资源

目标：注册并购买腾讯云 COS 服务，配置 SDK

步骤：

1.购买对象存储服务：https://console.cloud.tencent.com/cos

2.生成腾讯云密钥：https://console.cloud.tencent.com/cam/capi

2.COS SDK 开发文档： https://cloud.tencent.com/document/product/436/31215 

```shell
$ go get -u github.com/tencentyun/cos-go-sdk-v5
```



#### Redis(docker desktop)

why：提供邮箱验证码缓存功能

```shell
# 安装并启动一个redis容器
$ docker pull redis
$ docker run --name gredis -p 6379:6379 redis --requirepass "redisPassword"

# 进入容器 (cmd)
$ docker exec -it gredis bash
# 进入容器 (bash)
> redis-cli
# 登陆
127.0.0.1:6379> auth redisPassword
# 查看版本
127.0.0.1:6379> info
```

并将 `redisPassword` 设置为环境变量。
