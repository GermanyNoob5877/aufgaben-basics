package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	i := n
	counter := 0
	for ; i > 0; i-- {
		if n%i == 0 {
			counter++
		}
	}
	if counter > 2 {
		return false
	} else {
		return true
	}
}
