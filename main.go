package main

import (
	"flag"
	"log"

	"github.com/wangzhongyang/sql2proto/parser"
)

func parseFlag() *parser.Options {
	args := &parser.Options{}

	flag.StringVar(&args.InputFile, "f", "", "input file")
	flag.StringVar(&args.OutputDir, "o", "", "output dir")
	flag.StringVar(&args.Sql, "sql", "", "input SQL")

	flag.StringVar(&args.MysqlDsn, "db-dsn", "", "mysql dsn([user]:[pass]@/[database][?charset=xxx&...])")
	flag.StringVar(&args.MysqlTable, "db-table", "", "mysql table name")

	flag.Parse()
	return args
}

func main() {
	args := parseFlag()
	log.Println("generate begin...")
	if err := parser.Parser(args); err != nil {
		log.Fatalln("gen failed, err:", err)
	}
	log.Println("generate success...")
}
