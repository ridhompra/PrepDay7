package main

import "fmt"

func main() {
	data := []string{"Meja", "Buku", "Topi", "Baju", "Kayu"}
	fmt.Println(data)

	for i := 0; i < len(data)-1; i++ {
		if i == 0 {
			data = append([]string{"handuk"}, data...)
		}
	}
	data = append(data, "Celana")

	fmt.Println(data)
}
