package lib

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
