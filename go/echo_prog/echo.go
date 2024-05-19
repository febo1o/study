package main

import "fmt"

func main() {
	fmt.Println("Hello, user!")
	recurs("Hello, user!")
}

func recurs(x string) {
	fmt.Scan(&x)

	if (x != "exit" && x != "выход") {
		fmt.Println("Написано: " + x)
		recurs(x)
	}
}

