package main 

import "fmt"

func main() {

	arr1 := [10]int{};
	arr2 := [5]int{};

	responseMap1 := map[int]int{};
	responseMap2 := map[int]int{};

	fmt.Println("Preencha o primeiro vetor")
	for i := 0; i < len(arr1); i++ {
		fmt.Scan(&arr1[i]);
	}

	fmt.Println("Preencha o segundo vetor")
	for i := 0; i < len(arr2); i++ {
		fmt.Scan(&arr2[i]);
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] % 2 == 0 {
			
			acc := arr1[i];
			for j := 0; j < len(arr2); j++ {
				acc += arr2[j];
			}

			responseMap1[i] = acc;
		}
	}

	for i := 0; i < len(arr1); i++ {

		if arr1[i] % 2 != 0 {
			acc := arr1[i];
			for j := 0; j < len(arr2); j++ {
				acc += arr2[j];
			}

			responseMap2[i] = acc;
		}
	}

	for _, v := range responseMap1 {
		fmt.Printf("%d ", v);
	}
	fmt.Println("")
	for _, v := range responseMap2 {
		fmt.Printf("%d ", v);
	}
}