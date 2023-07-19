package main

import "fmt"

func main() {

	fmt.Println("DESAFIO #1")
	for i := 1; i != 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("###########################")
	fmt.Println("DESAFIO #2")

	for y := 1; y != 100; y++ {
		if y%3 == 0 {
			fmt.Println("Pin")
		}
		if y%5 == 0 {
			fmt.Println("Pan")
		}
		if y%5 != 0 && y%3 != 0 {
			fmt.Println(y)
		}
	}
}
