package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	fmt.Println(outer)
	for i := 0; i < (length - 2); i++ {
		fmt.Print(outer)
		for range i {
			fmt.Print(inner)
		}
		fmt.Println(outer)
	}
	for range length {
		fmt.Print(outer)
	}
	fmt.Println()
}
