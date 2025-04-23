package main

import "fmt"

func main() {
	arr := [10]int{};
	var isResponseFound bool = false;

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i]);
	}

	for i, v := range arr {
		if v > 50 {
			fmt.Printf("O número na posição %d é maior que 50\n", i);
			isResponseFound = true;
		}
	}

	if !isResponseFound {
		fmt.Println("Não há números maiores que 50 no array");
	}
}