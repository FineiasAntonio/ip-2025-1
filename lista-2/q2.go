package main

import "fmt"

func main() {
	var input int;

	fmt.Scan(&input);
	
	if input > 20 && input < 90 {
		fmt.Println("Verdadeiro");
	} else {
		fmt.Println("False");
	}

}