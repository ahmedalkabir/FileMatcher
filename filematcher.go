package filematcher

import (
	"crypto/md5"

	"github.com/ahmedalkabir/filematcher/matchers"
	"github.com/ahmedalkabir/filematcher/types"
)

type matcher struct {
	Detect matchers.MatcherFunc
	Type   types.TypesEnum
}

var listOfMatchers []matcher

func init() {
	listOfMatchers = []matcher{
		{matchers.Jpeg, types.JPEG_Type},
		{matchers.Png, types.PNG_Type},
		{matchers.Bmp, types.BMP_Type},
		{matchers.Pdf, types.PDF_Type},
		{matchers.Doc, types.DOC_Type},
		{matchers.Xls, types.XLS_Type},
		{matchers.Ppt, types.PPT_Type},
		{matchers.Zip, types.ZIP_Type},
	}
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
	for _, matcher := range listOfMatchers {
		if matcher.Detect(buf) {
			file.Type = matcher.Type
		}
	}

	if file.Type == types.Unknown_Type {
		return file, types.ErrUnSupportedFile
	}

	return file, nil
}

func CheckSumMD5(buf []byte, len_buf int) [16]byte {
	if len_buf != 0 {
		if len(buf) > len_buf {
			buf = buf[:len_buf]
		}
	}
	return md5.Sum(buf)
}
