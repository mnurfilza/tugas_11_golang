package main

import (
	"belajargolang/tugas11"
	"fmt"
)

func main() {
	var input1 string
	var input2 string
	fmt.Print("input data 1 disini")
	fmt.Scanln(&input1)
	fmt.Print("input data 2 disini")
	fmt.Scanln(&input2)
	tugas11.Tugas11(input1, input2)

	// tugas11.Perbedaan()
}
