package boyermoore

// Index returns the first index substr found in the s.
// function should return same result as `strings.Index` function
func Index(s string, substr string) int {
	d := CalculateSlideTable(substr)
	return IndexWithTable(&d, s, substr)
}

// IndexWithTable returns the first index substr found in the s.
// It needs the slide information of substr
func IndexWithTable(d *[256]int, s string, substr string) int {
	lsub := len(substr)
	ls := len(s)
	// fmt.Println(ls, lsub)
	switch {
	case lsub == 0:
		return 0
	case lsub > ls:
		return -1
	case lsub == ls:
		if s == substr {
			return 0
		}
		return -1
	}

	i := 0
	for i+lsub-1 < ls {
		j := lsub - 1
		for ; j >= 0 && s[i+j] == substr[j]; j-- {
		}
		if j < 0 {
			return i
		}

		slid := j - d[s[i+j]]
		if slid < 1 {
			slid = 1
		}
		i += slid
	}
	return -1
}

// CalculateSlideTable builds sliding amount per each unique byte in the substring
func CalculateSlideTable(substr string) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	for i := 0; i < len(substr); i++ {
		d[substr[i]] = i
	}
	return d
}
