package piscine

import "fmt"

func first_And_last_line_quadA(x int) {
	fmt.Print("o")
	for j := 0; j < x-2; j++ {
		fmt.Print("-")
	}
	fmt.Print("o")
	fmt.Println()
}

// [x:5,y:3]
func QuadA(x, y int) {
	if y > 1 && x > 2 { // implementation de #1

		first_And_last_line_quadA(x)
		defer first_And_last_line_quadA(x)
		for i := 0; i < y-2; i++ {
			fmt.Print("|")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("|")
			fmt.Println()
		}

	} else if x > 2 && y == 1 { // implementation de Program #2     o---o
		first_And_last_line_quadA(x)
	} else if x == 1 && y == 1 { // implementation de Program #3    o
		fmt.Println("o")
	} else if x == 1 && y > 2 { //  impementation de Program #4
		fmt.Println("o")
		for i := 0; i < y-2; i++ {
			fmt.Println("|")
		}
		fmt.Println("o")
	} else {
		fmt.Println("o")
		fmt.Println("o")
	}
}

/*

piscine.QuadA(5,3)
fmt.Println()

piscine.QuadA(5,1)
fmt.Println()

piscine.QuadA(1,1)
fmt.Println()

piscine.QuadA(1,5)
fmt.Println()

piscine.QuadA(5,5)
fmt.Println()

*/
