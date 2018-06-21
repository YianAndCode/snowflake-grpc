# Snowflake gRPC

## 介绍
本项目为 Snowflake 算法在 gRPC 中的实现，server 采用 Golang 编写，示例 client 采用 PHP 编写(`client/`)。

## 使用方法（Server）

### 编译

直接执行
```bash
./build.sh
```
编译后的可执行程序将会输出到 `bin/`

### 配置

在可执行程序同一目录一下保存一份名为 `config.json` 的配置文件，示例：

```JSON
{
    "server": "127.0.0.1",
    "nodeId": 1,
    "port": 6666
}
```

*也可以直接复制 `config.json.example` 到 `bin/` 并改名为 `config.json`*

### 运行

```bash
cd bin
./snowflake-grpc-server
```

## 使用方法（Client）

### 安装依赖

通过 composer 安装依赖和编译自动加载

```bash
cd client
composer install && composer dumpautoload
```

### 运行

```bash
php client.php
```