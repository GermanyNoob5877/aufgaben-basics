package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	for i := 0; i < width; i++ {
		fmt.Print("#")
	}
	fmt.Println()

	// Mittelteil
	for row := 0; row < height-2; row++ {
		fmt.Print("#")
		for i := 0; i < width-2; i++ {
			fmt.Print("#")
		}
		fmt.Println("#") //Hier Zeilenumbruch für nächste Zeile
	}

	// untere Zeile
	if height > 1 {
		for i := 0; i < width; i++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
