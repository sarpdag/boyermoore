package boyermoore

// IndexString returns the first index substr found in the s.
// function should return same result as `strings.Index` function
func IndexString(s string, substr string) int {
	d := CalculateSlideTableString(substr)
	return IndexWithTableString(&d, s, substr)
}

// IndexWithTableString returns the first index substr found in the s.
// It needs the slide information of substr calculated from CalculateSlideTableString.
func IndexWithTableString(d *[256]int, s string, substr string) int {
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

// CalculateSlideTableString builds sliding amount per each unique byte in the substring
func CalculateSlideTableString(substr string) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	for i := 0; i < len(substr); i++ {
		d[substr[i]] = i
	}
	return d
}

// Index returns the first index substr found in the s.
// function should return same result as `strings.Index` function
func Index(s, substr []byte) int {
	d := CalculateSlideTable(substr)
	return IndexWithTable(&d, s, substr)
}

// CalculateSlideTable builds sliding amount per each unique byte in the substring
func CalculateSlideTable(substr []byte) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	for i := 0; i < len(substr); i++ {
		d[substr[i]] = i
	}
	return d
}

// IndexWithTable returns the first index substr found in the s.
// It needs the slide information of substr calculated from CalculateSlideTable.
func IndexWithTable(d *[256]int, s, substr []byte) int {
	lsub := len(substr)
	ls := len(s)
	// fmt.Println(ls, lsub)
	switch {
	case lsub == 0:
		return 0
	case lsub > ls:
		return -1
	case lsub == ls:
		if byteSlicesEqual(s, substr) {
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

func byteSlicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if x != b[i] {
			return false
		}
	}

	return true
}

// IndexRev returns the last index substr found in the s.
func IndexRev(s, substr []byte) int {
	d := CalculateSlideTableRev(substr)
	return IndexWithTableRev(&d, s, substr)
}

// CalculateSlideTableRev builds sliding amount per each unique byte in the substring for a reverse search
func CalculateSlideTableRev(substr []byte) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	l := len(substr)
	for i := 0; i < len(substr); i++ {
		d[substr[l-i-1]] = i
	}
	return d
}

// IndexWithTableRev returns the last index substr found in the s.
// It needs the slide information of substr calculated from CalculateSlideTableRev.
func IndexWithTableRev(d *[256]int, s, substr []byte) int {
	lsub := len(substr)
	ls := len(s)
	// fmt.Println(ls, lsub)
	switch {
	case lsub == 0:
		return ls - 1
	case lsub > ls:
		return -1
	case lsub == ls:
		if byteSlicesEqual(s, substr) {
			return 0
		}
		return -1
	}

	i := 0
	for i+lsub-1 < ls {
		j := lsub - 1
		for ; j >= 0 && s[ls-(i+j)-1] == substr[lsub-j-1]; j-- {
		}
		if j < 0 {
			return ls - i - lsub
		}

		slid := j - d[s[ls-(i+j)-1]]
		if slid < 1 {
			slid = 1
		}
		i += slid
	}
	return -1
}

// IndexRev returns the last index substr found in the s.
func IndexRevString(s, substr string) int {
	d := CalculateSlideTableRevString(substr)
	return IndexWithTableRevString(&d, s, substr)
}

// CalculateSlideTableRev builds sliding amount per each unique byte in the substring for a reverse search
func CalculateSlideTableRevString(substr string) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	l := len(substr)
	for i := 0; i < len(substr); i++ {
		d[substr[l-i-1]] = i
	}
	return d
}

// IndexWithTableRev returns the last index substr found in the s.
// It needs the slide information of substr calculated from CalculateSlideTableRevString.
func IndexWithTableRevString(d *[256]int, s, substr string) int {
	lsub := len(substr)
	ls := len(s)
	// fmt.Println(ls, lsub)
	switch {
	case lsub == 0:
		return ls - 1
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
		for ; j >= 0 && s[ls-(i+j)-1] == substr[lsub-j-1]; j-- {
		}
		if j < 0 {
			return ls - i - lsub
		}

		slid := j - d[s[ls-(i+j)-1]]
		if slid < 1 {
			slid = 1
		}
		i += slid
	}
	return -1
}
