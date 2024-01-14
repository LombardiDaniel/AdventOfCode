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

func NavigateSantaFloors(s string) int {

	var floor int = 0

	// fmt.Println(s)

	for i, ch := range s {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}

		// fmt.Println(string(ch))
		// fmt.Println(floor)

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}