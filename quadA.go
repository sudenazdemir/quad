package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		// Bu satır sayısını belirler
		for i := 1; i <= y; i++ {
			// Bu koşul ilk ve son satırı belirtmek
			if i == 1 || i == y {
				// Bu döngü ilk ve son satırda yazılacak kodu belirler
				for j := 1; j <= x; j++ {
					// Bu koşul ilk ve son satırındaki ilk ve son simgeyi belirler
					if j == 1 || j == x {
						z01.PrintRune('o')
						// Bu ise ilk ve son satırın geri kalan kısmındaki simgeyi belirler
					} else if j > 1 && j < x {
						z01.PrintRune('-')
					}
				}
				// Bu koşul ilk ve son satır dışındaki satırları belirtir
			} else {
				// Bu  döngü ilk ve son satır dışındaki satırlarda yazılacak kodu belirler
				for j := 1; j <= x; j++ {
					// Bu koşul aradaki satırların ilk ve son simgesini belirler
					if j == 1 || j == x {
						z01.PrintRune('|')
						// Bu koşul ise aradaki satırların geri kalan kısmındaki simgeyi belirler
					} else if j > 1 && j < x {
						z01.PrintRune(' ')
					}
				}
			}
			// Burası her satır bittiğinde bir sonraki satıra geçer
			z01.PrintRune('\n')
		}
	}
}
