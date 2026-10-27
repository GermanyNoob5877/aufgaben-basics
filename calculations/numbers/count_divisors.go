package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	i := n
	counter := 0
	for ; i > 0; i-- {
		if n%i == 0 {
			counter++
		}
	}
	return counter
}
