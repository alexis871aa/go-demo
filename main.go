package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			//fmt.Println("Не заданы параметры для расчёта")
			//continue
			panic("Не заданы параметры для расчёта") // используется для критически важных ошибок, из-за которых рушится всё приложение
		}
		outputResult(IMT)
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать ещё расчёт (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
