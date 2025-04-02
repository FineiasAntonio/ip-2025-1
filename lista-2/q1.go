package main

import "fmt"

func main() {
	var input int32;

	fmt.Scan(&input);

	if input > 0 {
		fmt.Println("Positivo");
	} else if input < 0 {
		fmt.Println("Negativo");
	} else {
		fmt.Println("Nulo");
	}

}