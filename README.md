## What is sql2proto
It is common to design the database first, and then write a proto file, which may contain objects with a table structure.
If you don't find a relevant program on the Internet, you'll just pick one up

[中文说明](./README-zh.md)

## sql2proto
There are three main features:
1. Generate proto file from SQL file
2. Generate proto file based on MySQL dsn and table name (single or multiple table names split with ',')
3. Generate proto file from SQL statement

## Getting started
### Environments
```shell
go version
go version go1.18.10 darwin/amd64
` ` `
### Installation
```shell
go build github.com/wangzhongyang/sql2proto
` ` `
### Use
```shell
# Read SQL file to generate proto file
./sql2proto -o=/Users/wangzhongyang/go/src/sql2proto/example  -f=/Users/wangzhongyang/go/src/sql2proto/example/database.sql

# Connect to database to generate proto file
./sql2proto -o=/Users/wangzhongyang/go/src/sql2proto/example  -db-dsn="sql2proto:sql2proto@tcp(zhongyang.wang:3306)/sql2proto? charset=utf8mb4&parseTime=True&loc=Local" -db-table="test_table,test_table2"
```

***proto generates go files ***
```shell
protoc --proto_path=./example --go_out=. ./example/gen.proto
```

## License
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