package main

import "fmt"

var cnt int = 0;


func main() {
	fmt.Println("Hello, enter your username!")
	recurs("Hello, enter your username!")
}

func recurs(x string) {
	fmt.Scan(&x)
	if (x != "exit" && x != "выход") {
		if (cnt == 0) {
			cnt += 1
			fmt.Println("Hello, " + x + "!")
			recurs(x)
		} else {
			fmt.Println("Написано: " + x)
			recurs(x)
		}
	}
}

