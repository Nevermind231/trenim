package main

import (
	"fmt"
	"math"
	"time"
)

const IMTpower = 2

func main() {
	fmt.Println("---Калькулятор индекса массы тела---")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
	fmt.Println("")

	if IMT > 26 {
		fmt.Println("Худей жирнич")
	}
	if IMT >= 18 {
		fmt.Println("Нормальный сухарь")
	} else {
		fmt.Println("Тебе бы похавать")
	}

	time.Sleep(5 * time.Second)

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела %.0f", IMT)
	fmt.Print(result)

}

func calculateIMT(userKg, userHeight float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTpower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой вес")
	fmt.Scan(&userKg)
	fmt.Print("Введите свой рост в сантиметрах")
	fmt.Scan(&userHeight)
	return userKg, userHeight
}
