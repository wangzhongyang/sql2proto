package parser

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type ReadFromMySQL struct {
	options *Options
}

func (r *ReadFromMySQL) GetStructList() ([]*StructInfo, error) {
	if len(r.options.MysqlTable) == 0 {
		return nil, fmt.Errorf("table's number failed")
	}
	content := ""
	db, err := gorm.Open(mysql.Open(r.options.MysqlDsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open MySQL failed, err:%w", err)
	}
	tables := strings.Split(r.options.MysqlTable, ",")
	for _, tableName := range tables {
		sql, err := r.getCreateTableSQL(db, tableName)
		if err != nil {
			return nil, err
		}
		content += sql + ";\n"
	}
	structArr, err := ParseSql(content)
	if err != nil {
		return nil, err
	}

	return structArr, nil
}

func (r *ReadFromMySQL) getCreateTableSQL(db *gorm.DB, tableName string) (string, error) {
	var showTable ShowTable
	sql := fmt.Sprintf("SHOW CREATE TABLE %s", tableName)
	if err := db.Raw(sql).Scan(&showTable).Error; err != nil {
		return "", fmt.Errorf("get create table sql failed, sql:%s, err:%w", sql, err)
	}
	return showTable.CreateTable, nil
}

type ShowTable struct {
	Table       string `gorm:"Table"`
	CreateTable string `gorm:"column:Create Table"`
}
