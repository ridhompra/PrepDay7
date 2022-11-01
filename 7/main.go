package main

import "fmt"

func main() {
	var input int
	var hasil []int
	fmt.Print("Masukan limit :")
	fmt.Scan(&input)
	for i := 1; i <= input; i++ {
		if i%2 == 1 {
			hasil = append(hasil, i)
		}
	}
	fmt.Printf("Bilangan fanjil dari %d adalah . . .\n %d\n", input, hasil)
}
