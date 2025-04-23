package main

import "fmt"

func main() {
	arr := [10]int{};

	fmt.Println("Preencha o vetor");
	
	for i, _ := range arr {
		fmt.Scan(&arr[i]);
	}

	fmt.Println("Números pares");
	for _, v := range arr {
		if v % 2 == 0 {
			fmt.Printf("%d ", v);
		}
	}

	acc := 0;
	for _, v := range arr {
		if v % 2 == 0 {
			acc += v;
		}
	}
	fmt.Printf("\nSoma dos números pares: %d\n", acc);
	
	fmt.Println("Números ímpares");
	for _, v := range arr {
		if v % 2 != 0 {
			fmt.Printf("%d ", v);
		}
	}
}