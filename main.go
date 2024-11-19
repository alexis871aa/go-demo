package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64 // по-умолчанию будет 0.0
	var userKg float64     // тоже по-умолчанию будет 0.0
	fmt.Println("___ Калькулятор индекса массы тела ___")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}
