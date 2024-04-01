package matchers

// Sources
// https://gist.github.com/leommoore/f9e57ba2aa4bf197ebc5
// https://en.wikipedia.org/wiki/List_of_file_signatures
// Doc: https://bz.apache.org/ooo/show_bug.cgi?id=111457

type MatcherFunc func(buf []byte) bool

// Images ============================================================================
func Jpeg(buf []byte) bool {
	if len(buf) > 3 {
		return buf[0] == 0xFF &&
			buf[1] == 0xD8 &&
			buf[2] == 0xFF &&
			buf[3] == 0xE0
	}
	return false
}

func Png(buf []byte) bool {
	if len(buf) > 3 {
		// it's ascii digit if you look
		// well .PNG
		return buf[0] == 0x89 &&
			buf[1] == 0x50 &&
			buf[2] == 0x4e &&
			buf[3] == 0x47
	}
	return false
}

func Bmp(buf []byte) bool {
	if len(buf) > 1 {
		return buf[0] == 0x42 &&
			buf[1] == 0x4D
	}
	return false
}

func Psd(buf []byte) bool {
	if len(buf) > 3 {
		return buf[0] == 0x38 &&
			buf[1] == 0x42 &&
			buf[2] == 0x50 &&
			buf[3] == 0x53
	}
	return false
}

// Documents ======================================================================
func Pdf(buf []byte) bool {
	// for the context
	// based on the PDF Specification
	// the first line of a PDF file shall be a header consisting of the 5 characters %PDF-
	// followed by a version number of the form 1.N where N is a digit between 0 and 7.

	// A conforming reader shall accept files with any of the following headers:
	// %PDF-1.0
	// %PDF-1.1
	// etc....

	// but there's some files don't follow these requirements proposed by PDF Specifications
	// and put the header at first 1024 bytes instead of the first line of a PDF

	buf_length := len(buf)
	pdfHeader := []byte{0x25, 0x50, 0x44, 0x46, 0x2D}

	if buf_length > 4 {

		// a valid correct pdf
		if compareSlices(buf[:5], pdfHeader, 5) {
			return true
		} else {
			// peek the slices until we find the header
			for i := range buf {
				if compareSlices(buf[i:i+5], pdfHeader, 5) {
					return true
				}
			}
		}
	}
	return false
}

// for Office Documents it's kinda different from other files so it need to be handled
// in different ways
func Doc(buf []byte) bool {
	if len(buf) > 513 {
		return buf[0] == 0xD0 && buf[1] == 0xCF &&
			buf[2] == 0x11 && buf[3] == 0xE0 &&
			buf[512] == 0xEC && buf[513] == 0xA5
	}
	return false
}

func Xls(buf []byte) bool {
	if len(buf) > 513 {
		isItOfficeFileFirst := buf[0] == 0xD0 &&
			buf[1] == 0xCF &&
			buf[2] == 0x11 &&
			buf[3] == 0xE0

		switch {
		case isItOfficeFileFirst && buf[512] == 0x09 && buf[513] == 0x08:
			return true
		case isItOfficeFileFirst && buf[512] == 0xFD && buf[513] == 0xFF:
			return true
		}
	}
	return false
}

func Ppt(buf []byte) bool {
	if len(buf) > 513 {
		return buf[0] == 0xD0 && buf[1] == 0xCF &&
			buf[2] == 0x11 && buf[3] == 0xE0 &&
			buf[512] == 0xA0 && buf[513] == 0x46 // to make sure it's a really Ppt file
	}
	return false
}

func Zip(buf []byte) bool {
	if len(buf) > 3 {
		isItZipFirst := buf[0] == 0x50 && buf[1] == 0x4B

		switch {
		case isItZipFirst && buf[2] == 0x03 && buf[3] == 0x04:
			return true
		case isItZipFirst && buf[2] == 0x05 && buf[3] == 0x06: // empty archive
			return true
		case isItZipFirst && buf[2] == 0x07 && buf[3] == 0x08: // spanned archive
			return true
		}
	}
	return false

}
