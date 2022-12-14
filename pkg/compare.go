package pkg

func CompareIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func CheckMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
