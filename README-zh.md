## 起因
工作中一般都是先设计数据库，然后写proto文件，其中proto文件中可能会包含表结构的对象。

因为在网上没有找到相关的程序，就开发了这个程序

***如果觉得有用，欢迎给个star***

## sql2proto
目前主要包含三个功能：
1. 根据SQL文件生成proto文件
2. 根据MySQL的dsn和表名（单个或用','分割的多个表名）生成proto文件
3. 根据SQL语句生成proto文件

## 使用
### 环境
```shell
go version
go version go1.18.10 darwin/amd64
```
### 安装
```shell
go build github.com/wangzhongyang/sql2proto
```
### 使用
```shell
# 读取SQL文件生成proto文件
./sql2proto -o=/Users/wangzhongyang/go/src/sql2proto/example -f=/Users/wangzhongyang/go/src/sql2proto/example/database.sql

# 连接数据库生成proto文件
./sql2proto -o=/Users/wangzhongyang/go/src/sql2proto/example -db-dsn="sql2proto:sql2proto@tcp(zhongyang.wang:3306)/sql2proto?charset=utf8mb4&parseTime=True&loc=Local" -db-table="test_table,test_table2"
```

***proto生成go文件***
```shell
protoc --proto_path=./example --go_out=. ./example/gen.proto
```


## 开源声明
Copyright [2023] [wangzhongyang](https://github.com/wangzhongyang)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.