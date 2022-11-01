package main

import "fmt"

func main() {
	data := []string{"Jeruk", "Pepaya", "Jambu", "Anggur", "Melon"}
	fmt.Println(data)
	for i := 0; i < len(data)-1; i++ {
		if data[i] == "Jambu" {
			data = append(data[:i], data[i+1:]...)
		}
	}
	fmt.Println(data)
}
