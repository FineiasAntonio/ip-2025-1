package main

import "fmt"

func main() {
	sum := 0;
	qnt := 0;

	for i := 50; i <= 70; i += 2 {
		sum += i;
		qnt++;
	}

	fmt.Printf("Soma: %d\n", sum);
	fmt.Printf("Média: %d\n", sum /qnt);
}