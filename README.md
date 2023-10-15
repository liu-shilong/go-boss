# Go-Boss
基于Gin开发的多应用后台管理系统

## 1.目录结构
<details>
<summary>展开查看</summary>
<pre><code>
├── api API接口
│   ├── http HTTP接口
│   │   ├── v1
│   ├── rpc RPC接口
│   │   ├── v1
├── cmd 命令行
├── internal 业务逻辑
│   ├── middleware 中间件
│   ├── model 模型
│   ├── service 服务
├── lang 多语言
├── pkg 扩展包
│   ├── auth 权限
│   ├── captcha 验证码
│   ├── crypt 加密
│   ├── database 数据库处理
│   │   ├── etcd
│   │   ├── etcd
│   │   ├── mongo
│   ├── datetime 日期处理
│   ├── file 文件
│   ├── http 请求响应
│   │   ├── request 请求
│   │   ├── response 响应
│   ├── logger 日志
│   ├── mq 消息队列
│   ├── platform 第三方平台
│   │   ├── alibaba 阿里巴巴
│   │   │   ├── alipay 支付宝
│   │   │   ├── dingding 钉钉
│   │   ├── bytedance 字节跳动
│   │   │   ├── douyin 抖音
│   │   │   ├── feishu 飞书
│   │   ├── tencent 腾讯
│   │   │   ├── mp 小程序
│   │   │   ├── wechat 微信
│   │   │   ├── wework 企业微信
│   ├── qrcode 二维码
│   ├── search-engine 搜索引擎
│   ├── trace 追踪
│   ├── util 通用
├── router 路由
├── runtime 运行时
│   ├── logs 日志
│   │   ├── app 应用日志
│   │   ├── sql 数据库日志
└── test 测试
</code></pre>
</details>


## 2.启动

### 2.1 创建数据库
```
go_boss
```
### 2.2 执行运行命令
```
go run main.go
```
或
```
go build -o go-boss.exe main.go start
```
```
./go-boss start
```

## 3. 运行效果

### 3.1 测试url
#### 3.1.1 请求
```
http://localhost:8080/v1/ping
```
#### 3.1.2 响应
```
{
  "code":200,
  "msg":"pong"
}
```
### 3.2 业务url
#### 3.2.1执行sql
```
INSERT INTO `go_boss`.`admin` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `mobile`, `password`) VALUES (1, NULL, NULL, NULL, 'admin', '13866668888', '');
```
#### 3.2.2 请求
```
http://localhost:8080/v1/admin
```
#### 3.2.3 响应
```
{
  "code":200,
  "msg":"success",
  "data":{
  "mobile":"13866668888",
  "name":"admin"
  }
}
```




