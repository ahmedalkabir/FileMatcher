package types

import (
	"errors"
)

type TypesEnum int

const (
	Unknown_Type TypesEnum = iota + 1
	JPEG_Type
	JPG_Type
	PNG_Type
	BMP_Type
	PDF_Type
	DOCX_Type
	DOC_Type
	XLSX_Type
	XLS_Type
	PPTX_Type
	PPT_Type
	ZIP_Type
)

func (i TypesEnum) String() string {
	return [...]string{"Unknown Type", "jpeg", "jpg", "png",
		"BMP", "pdf", "docx", "doc", "xlsx", "xls", "pptx", "ppt", "zip"}[i-1]
}

type FileMatcher struct {
	Type TypesEnum
}

// Errors =======================================================================

// UnSupportedFile
var ErrUnSupportedFile = errors.New("Error: The File is Not Valid or Corrupted!!")

// UnSupportedFile
var ErrEmptyBuffer = errors.New("Error: Buffer is Empty")
