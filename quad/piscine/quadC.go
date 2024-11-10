package piscine

import "fmt"

func first_line_quadC(x int) {
	for i := 0; i <= 0; i++ {
		fmt.Print("A")
		for j := 0; j < x-2; j++ {
			fmt.Print("B")
		}
		fmt.Print("A")
		fmt.Println()
	}
}

func last_line_quadC(x int) {
	for i := 0; i <= 0; i++ {
		fmt.Print("C")
		for j := 0; j < x-2; j++ {
			fmt.Print("B")
		}
		fmt.Print("C")
		fmt.Println()
	}
}

// [x:5,y:3]
func QuadC(x, y int) {
	if y > 1 && x > 2 { // implementation de #1

		first_line_quadC(x)
		for i := 0; i < y-2; i++ {
			fmt.Print("B")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("B")
			fmt.Println()
		}
		last_line_quadC(x)

	} else if x > 2 && y == 1 { // implementation de Program #2     /***\
		first_line_quadC(x)
	} else if x == 1 && y == 1 { // implementation de Program #3    /
		fmt.Println("A")
	} else if x == 1 && y > 2 { //  impementation de Program #4
		fmt.Println("A")
		for i := 0; i < y-2; i++ {
			fmt.Println("B")
		}
		fmt.Println("C")
	} else {
		fmt.Println("A")
		fmt.Println("C")
	}
}

// for test cases try the following :

/*
piscine.QuadC(5,3)
fmt.Println()

piscine.QuadC(5,1)
fmt.Println()

piscine.QuadC(1,1)
fmt.Println()

piscine.QuadC(1,5)
fmt.Println()

piscine.QuadC(5,5)
fmt.Println()

*/
