package file

type FileData struct {
	Columns []string `json:"columns"`
	Table_Data [][]interface{} `json:"table_data"`
}