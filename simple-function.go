package main

import "fmt"

func main() {
	listNum := []float64{5, 4, 3, 2, 1}
	fmt.Println(listNum[:5])
	fmt.Println("Sum: ", add(listNum))
}

func add(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
