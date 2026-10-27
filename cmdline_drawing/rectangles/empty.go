package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	// obere Zeile
	for i := 0; i < width; i++ {
		fmt.Print("#")
	}
	fmt.Println()

	// Mittelteil
	for row := 0; row < height-2; row++ {
		fmt.Print("#")
		for i := 0; i < width-2; i++ {
			fmt.Print(" ")
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

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
