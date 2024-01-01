package filematcher

import (
	"errors"

	filematcher "github.com/ahmedalkabir/filematcher/types"
)

// func Match(buf []byte) (FileMatcher, error) {
// 	var file FileMatcher

// 	length := len(buf)

// 	if length == 0 {
// 		return "", errors.New("EmptyBuffer")
// 	}

// 	return "Type", nil
// }

// func Match(buf []byte) (FileMatcher, error) {
// 	var file FileMatcher

// 	length := len(buf)

// 	if length == 0 {
// 		return file, errors.New("EmptyBuffer")
// 	}

// 	file.Type = DOCX_Type
// 	return file, nil
// }

func Match(buf []byte) (filematcher.FileMatcher, error) {
	var file filematcher.FileMatcher

	length := len(buf)

	if length == 0 {
		return file, errors.New("EmptyBuffer")
	}

	file.Type = filematcher.DOCX_Type
	return file, nil
}
