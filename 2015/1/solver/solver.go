package solver

// Counts the number of that char present in the string.
// Returns int
func CountCharInStr(c rune, s string) int {

	var count int = 0
	for _, ch := range s {
		if c == ch {
			count++
		}
	}

	return count
}

func CountSantaFloors(s string) int {

	up := CountCharInStr('(', s)
	down := CountCharInStr(')', s)

	return up - down
}