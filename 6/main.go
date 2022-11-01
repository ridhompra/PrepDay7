package main

import "fmt"

func main() {
	var input1, input2 int
	var result []int
	fmt.Print("Input 1\t: ")
	fmt.Scan(&input1)
	fmt.Print("Input 2\t: ")
	fmt.Scan(&input2)
	for i := input1; i <= input2; i++ {
		result = append(result, i)
	}
	fmt.Println("\noutput : ", result)
}
