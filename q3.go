package main

import "fmt"

func main() {

	acc := 0;

	for i := 0; i <= 100; i++ {
		fmt.Println(i);
		acc += i;
	}

	fmt.Printf("Soma: %d\n", acc);

}