package numbers

// Berechnet das Minimum von zwei Zahlen.
func Min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

// Berechnet das Maximum von zwei Zahlen.
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
