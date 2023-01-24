package test

import (
	"sql2proto/parser"
	"testing"
)

func Test_ReadFromFile_GetStructList(t *testing.T) {
	options := &parser.Options{
		InputFile: "/Users/wangzhongyang/go/src/sql2proto/example/database.sql",
		OutputDir: "/Users/wangzhongyang/go/src/sql2proto/example",
	}
	if err := parser.Parser(options); err != nil {
		t.Fatal(err)
	}
}

func Test_ReadFromMySQL_GetStructList(t *testing.T) {
	dns := "sql2proto:sql2proto@tcp(zhongyang.wang:3306)/sql2proto?charset=utf8mb4&parseTime=True&loc=Local"
	tables := "test_table,test_table2"
	options := &parser.Options{
		OutputDir:  "/Users/wangzhongyang/go/src/sql2proto/example",
		MysqlDsn:   dns,
		MysqlTable: tables,
	}
	if err := parser.Parser(options); err != nil {
		t.Fatal(err)
	}
}
