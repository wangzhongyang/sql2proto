package test

import (
	"fmt"
	"os"
	"sql2proto/parser"
	"testing"
)

var dirPath = ""

func TestMain(m *testing.M) {
	dirPath = fmt.Sprintf("%s/src", os.Getenv("GOPATH"))
	m.Run()
}

func Test_ReadFromFile_GetStructList(t *testing.T) {
	options := &parser.Options{
		InputFile: fmt.Sprintf("%s/sql2proto/example/database.sql", dirPath),
	}
	if err := parser.Parser(options); err != nil {
		t.Fatal(err)
	}
}

func Test_ReadFromMySQL_GetStructList(t *testing.T) {
	dns := "sql2proto:sql2proto@tcp(zhongyang.wang:3306)/sql2proto?charset=utf8mb4&parseTime=True&loc=Local"
	tables := "test_table,test_table2"
	options := &parser.Options{
		OutputDir:  fmt.Sprintf("%s/sql2proto/example", dirPath),
		MysqlDsn:   dns,
		MysqlTable: tables,
	}
	if err := parser.Parser(options); err != nil {
		t.Fatal(err)
	}
}

func Test_ReadFromSQL_GetStructList(t *testing.T) {
	filePath := fmt.Sprintf("%s/sql2proto/example/database.sql", dirPath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("open file faield, path:%s, err:%v", filePath, err)
	}
	options := &parser.Options{
		Sql: string(content),
	}
	protoStr, err := parser.ParserToString(options)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(protoStr)
	t.Log(protoStr)
}
