package lib

import "strings"

func IsContainQ(target string) bool {
	target = strings.ToLower(target)
	for i := 0; i < len(target); i++ {
		if target[i] == 'q' {
			return true
		} else if target[i] > 'q' {
			break
		}
	}
	return false
}

func IsContainU(target string) bool {
	target = strings.ToLower(target)
	for i := 0; i < len(target); i++ {
		if target[i] == 'u' {
			return true
		} else if target[i] > 'u' {
			break
		}
	}
	return false
}

func IsQUCase(word string) bool {
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		if word[i] == 'q' && i+1 < len(word) {
			return word[i+1] == 'u'
		}
	}
	return false
}

func IsDeductCase(word string, u bool, q bool) bool {
	if !q {
		return false
	} else if q && u {
		return false
	}

	return IsContainQ(word) && IsContainU(word)
}

func CountQU(sorted string) (int, int) {
	q := 0
	u := 0
	for i := 0; i < len(sorted); i++ {
		if sorted[i] > 'u' {
			break
		} else if sorted[i] == 'q' {
			q++
		} else if sorted[i] == 'u' {
			u++
		}
	}

	return q, u
}
