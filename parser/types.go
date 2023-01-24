package parser

type FileInfo struct {
	FilePath   string
	GenPath    string
	DirName    string
	StructList []*StructInfo
}

type StructInfo struct {
	Name   string
	Fields []StructInfoField
}

type StructInfoField struct {
	FieldName    string
	FieldType    string
	FieldComment string
}

type Options struct {
	InputFile string
	OutputDir string
	Sql       string

	MysqlDsn   string
	MysqlTable string
}

type Reader interface {
	GetStructList() ([]*StructInfo, error)
}
