package rectangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein Quadrat mit dieser Seitenlänge auf der Konsole.
// Das Quadrat soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidSquare(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("#")
	}
	fmt.Println()

	// Mittelteil
	for row := 0; row < length-2; row++ {
		fmt.Print("#")
		for i := 0; i < length-2; i++ {
			fmt.Print("#")
		}
		fmt.Println("#") //Hier Zeilenumbruch für nächste Zeile
	}

	// untere Zeile
	if length > 1 {
		for i := 0; i < length; i++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
