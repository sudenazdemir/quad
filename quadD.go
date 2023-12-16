package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		// Bu döngü satırları oluşturur
		for i := 1; i <= y; i++ {
			// ilk ve son satırı belirtir
			if i == 1 || i == y {
				// ilk satırdaki kodu belirtir
				for j := 1; j <= x; j++ {
					// ilk satırdaki ilk simgeyi belirtir
					if j == 1 {
						z01.PrintRune('A')
						// ilk satırdaki aradaki simgeleri belirtir
					} else if j > 1 && j < x {
						z01.PrintRune('B')
						// ilk satırdaki son simgeyi belirtir
					} else if j == x {
						z01.PrintRune('C')
					}
				}

				// aradaki satırı belirtir
			} else if i > 1 && i < y {
				// aradaki satırlarda olan kodu belirtir
				for j := 1; j <= x; j++ {
					// ara satırlardaki ilk ve son simgeyi belirtir
					if j == 1 || j == x {
						z01.PrintRune('B')
						// aradaki satırların geri kalan kısmındaki simgeyi belirler
					} else if j > 1 && j < x {
						z01.PrintRune(' ')
					}
				}
			}
			// Bir sonraki satıra geçer
			z01.PrintRune('\n')
		}
	}
}
