package filematcher

type TypesEnum int

const (
	JPEG_Type TypesEnum = iota + 1
	JPG_Type
	PNG_Type
	PDF_Type
	DOCX_Type
	DOC_Type
	XLSX_Type
	XLS_Type
)

func (i TypesEnum) String() string {
	return [...]string{"jpeg", "jpg", "png",
		"pdf", "docx", "doc", "xlsx", "xls"}[i-1]
}

type FileMatcher struct {
	Type TypesEnum
}
