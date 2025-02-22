package main

import (
	"michail-lab/internal/input"
	"fmt"
	"slices"
)

func main() {
	fmt.Print("Размер массива> ")
	var arrSize uint
	_, err := fmt.Scanln(&arrSize)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	if arrSize < 0 {
		fmt.Println("Размер массива должен быть положительным")
		return
	}

	resultArr, err := input.Input(arrSize)
	if err != nil {
		fmt.Println("Ошибка ввода: ", err.Error())
		return
	}

	fmt.Println("Максимальное значение: ", slices.Max(resultArr))

}
