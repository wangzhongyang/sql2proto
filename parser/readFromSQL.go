package parser

type ReadFromSQL struct {
	options *Options
}

func (r *ReadFromSQL) GetStructList() ([]*StructInfo, error) {
	structArr, err := ParseSql(r.options.Sql)
	if err != nil {
		return nil, err
	}

	return structArr, nil
}
