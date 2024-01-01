package filematcher

import (
	"github.com/ahmedalkabir/filematcher/matchers"
	"github.com/ahmedalkabir/filematcher/types"
)

type Matcher struct {
	Detect matchers.MatcherFunc
	Type   types.TypesEnum
}

var ListOfMatchers []Matcher

func init() {
	ListOfMatchers = []Matcher{
		{matchers.Jpeg, types.JPEG_Type},
		{matchers.Png, types.PNG_Type},
		{matchers.Bmp, types.BMP_Type},
		{matchers.Pdf, types.PDF_Type},
		{matchers.Doc, types.DOC_Type},
		{matchers.Xls, types.XLS_Type}}
}

func Match(buf []byte) (types.FileMatcher, error) {
	var file types.FileMatcher
	file.Type = types.Unknown_Type

	length := len(buf)

	if length == 0 {
		return file, types.ErrEmptyBuffer
	}

	// let's iterate over matchers and
	// detect the file
	for _, matcher := range ListOfMatchers {
		if matcher.Detect(buf) {
			file.Type = matcher.Type
		}
	}

	if file.Type == types.Unknown_Type {
		return file, types.ErrUnSupportedFile
	}

	return file, nil
}
