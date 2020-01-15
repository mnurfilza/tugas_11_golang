package tugas11

import (
	"fmt"
	"strconv"
)

func Tugas11(a, b string) {
	num1, err := strconv.Atoi(a)
	if err != nil {
		return
	}

	num2, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	fmt.Println("ini penjumlahan")
	jumlah := num1 + num2
	fmt.Println(num1, "+", num2, "=", jumlah)

	fmt.Println("ini pengurangan")
	kurang := num1 - num2
	fmt.Println(num1, "-", num2, "=", kurang)

	fmt.Println("ini pembagian")
	bagi := num1 / num2
	fmt.Println(num1, "/", num2, "=", bagi)

	fmt.Println("ini perkalian")
	kali := num1 * num2
	fmt.Println(num1, "x", num2, "=", kali)

}
