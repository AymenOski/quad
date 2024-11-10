package piscine

import "fmt"

func first_line_quadB(x int) {
	for i := 0; i <= 0; i++ {
		fmt.Print("/")
		for j := 0; j < x-2; j++ {
			fmt.Print("*")
		}
		fmt.Print("\\")
		fmt.Println()
	}
}

func last_line_quadB(x int) {
	for i := 0; i <= 0; i++ {
		fmt.Print("\\")
		for j := 0; j < x-2; j++ {
			fmt.Print("*")
		}
		fmt.Print("/")
		fmt.Println()
	}
}

// [x:5,y:3]
func QuadB(x, y int) {
	if y > 1 && x > 2 { // implementation de #1

		first_line_quadB(x)
		for i := 0; i < y-2; i++ {
			fmt.Print("*")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
			fmt.Println()
		}
		last_line_quadB(x)

	} else if x > 2 && y == 1 { // implementation de Program #2     /***\
		first_line_quadB(x)
	} else if x == 1 && y == 1 { // implementation de Program #3    /
		fmt.Println("/")
	} else if x == 1 && y > 2 { //  impementation de Program #4
		fmt.Println("/")
		for i := 0; i < y-2; i++ {
			fmt.Println("*")
		}
		fmt.Println("\\")
	} else {
		fmt.Println("/")
		fmt.Println("\\")
	}
}

// for test cases try the following :

/*
piscine.QuadB(5,3)
fmt.Println()

piscine.QuadB(5,1)
fmt.Println()

piscine.QuadB(1,1)
fmt.Println()

piscine.QuadB(1,5)
fmt.Println()

piscine.QuadB(5,5)
fmt.Println()

*/
