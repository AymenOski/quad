package piscine

import "fmt"

func first_And_last_line_quadD(x int) {
	for i := 0; i <= 0; i++ {
		fmt.Print("A")
		for j := 0; j < x-2; j++ {
			fmt.Print("B")
		}
		fmt.Print("C")
		fmt.Println()
	}
}

// [x:5,y:3]
func QuadD(x, y int) {
	if y > 1 && x > 2 { // implementation de #1

		first_And_last_line_quadD(x)
		for i := 0; i < y-2; i++ {
			fmt.Print("B")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("B")
			fmt.Println()
		}
		defer first_And_last_line_quadD(x)

	} else if x > 2 && y == 1 { // implementation de Program #2     o---o
		first_And_last_line_quadD(x)
	} else if x == 1 && y == 1 { // implementation de Program #3    o
		fmt.Println("A")
	} else if x == 1 && y > 2 { //  impementation de Program #4
		fmt.Println("A")
		for i := 0; i < y-2; i++ {
			fmt.Println("B")
		}
		fmt.Println("A")
	} else {
		fmt.Println("A")
		fmt.Println("A")
	}
}

/*

piscine.QuadD(5,3)
fmt.Println()

piscine.QuadD(5,1)
fmt.Println()

piscine.QuadD(1,1)
fmt.Println()

piscine.QuadD(1,5)
fmt.Println()

piscine.QuadD(5,5)
fmt.Println()

*/
