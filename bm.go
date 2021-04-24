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
		for ; j >= 0; j-- {
			if s[i+j] != substr[j] {
				if slidx := d[s[i+j]]; slidx > -100 {
					val := slidx
					if slidx <= 0 {
						val = 1
					}
					i += val
				} else {
					i += lsub
				}
				break
			}
		}

		if j == -1 {
			return i
		}
	}
	return -1
}

// CalculateSlideTable builds sliding amount per each unique byte in the substring
func CalculateSlideTable(substr string) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -100
	}
	for i := 0; i < len(substr); i++ {
		d[substr[i]] = len(substr) - i - 2
	}
	return d
}
