package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	var arrSize uint
	var scanner *bufio.Scanner

	// Проверяем, есть ли файл с входными данными
	if _, err := os.Stat("/app/input.txt"); err == nil {
		// Чтение данных из файла
		file, err := os.Open("/app/input.txt")
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		fmt.Println("Use standart input", err)
		// Интерактивный ввод
		scanner = bufio.NewScanner(os.Stdin)
	}

	// Чтение размера массива
	fmt.Print("Размер массива> ")
	if scanner.Scan() {
		_, err := fmt.Sscan(scanner.Text(), &arrSize)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
	}

	if arrSize < 0 {
		fmt.Println("Размер массива должен быть положительным")
		return
	}

	// Чтение элементов массива
	resultArr := make([]int, arrSize)
	for i := uint(0); i < arrSize; i++ {
		fmt.Printf("Элемент %d> ", i+1)
		if scanner.Scan() {
			_, err := fmt.Sscan(scanner.Text(), &resultArr[i])
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
		}
	}

	fmt.Println("Максимальное значение: ", slices.Max(resultArr))
}
