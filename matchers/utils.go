package matchers

func compareSlices(first, second []byte, length int) bool {

	if len(first) != len(second) {
		return false
	}

	for i := 0; i < length; i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
