package parser

import (
	"fmt"
	"os"
)

type ReadFromFile struct {
	options *Options
}

func (r *ReadFromFile) GetStructList() ([]*StructInfo, error) {
	content, err := os.ReadFile(r.options.InputFile)
	if err != nil {
		return nil, fmt.Errorf("open file faield, path:%s, err:%w", r.options.InputFile, err)
	}
	structArr, err := ParseSql(string(content))
	if err != nil {
		return nil, err
	}

	return structArr, nil
}
