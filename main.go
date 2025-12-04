package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Records struct {
	Name        string
	Try         int
	UserNumbers []int
	Success     int
}
type User struct {
	UserNumber   int
	Limit        int
	randomNumber int
	Try          int
}

func main() {

	var records Records
	var name string
	fmt.Println("Введите ваше Имя:")
	fmt.Scanln(&name)
	records.Name = name
	user := User{

		Limit: 5,
		Try:   0,
	}

	records.UserNumbers = []int{}

	user.startGame(&records)

	saveToFile(records, user)
}
func (u *User) genRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	u.randomNumber = rand.Intn(10) + 1
}

func (u *User) startGame(records *Records) {
	u.Limit = 5
	u.Try = 0
	u.genRandomNumber()

	fmt.Println("Угадай число от 1 до 10")
	for u.Try < u.Limit {
		fmt.Println("Введите ваше число")
		fmt.Scanln(&u.UserNumber)
		if u.UserNumber < 1 || u.UserNumber > 10 {
			fmt.Println("Ошибка ввода. Число должно быть от 1 до 10")
			continue
		}
		u.Try++
		records.UserNumbers = append(records.UserNumbers, u.UserNumber)

		if u.UserNumber < u.randomNumber {
			fmt.Println("Бери повыше")
		} else if u.UserNumber > u.randomNumber {
			fmt.Println("Маленько меньше")
		} else {
			fmt.Println("Красавчик ты угадал")
			records.Success = u.Try
			return
		}
		if u.Try < u.Limit {
			fmt.Printf(" осталось %d попыток. \n", u.Limit-u.Try)
		} else {
			fmt.Printf("У тебя кончились попытки,загаданное число было %d\n", u.randomNumber)

		}
	}
}
func saveToFile(records Records, u User) {
	file, err := os.OpenFile("records.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "Имя: %s Числа: ", records.Name)
	if err != nil {
		fmt.Println("Ошибка при записи имени и чисел в файл:", err)
		return
	}
	for _, num := range records.UserNumbers {
		_, err := fmt.Fprintf(file, "%d ", num)
		if err != nil {
			fmt.Println("Ошибка при записи числа в файл:", err)
			return
		}
	}
	if records.Success > 0 {
		_, err = fmt.Fprintf(file, "Угадал на попытке: %d Загаданное число: %d\n", records.Success, u.randomNumber)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	} else {
		_, err = fmt.Fprintf(file, "Не угадал число за %d попыток. Загаданное число: %d\n", len(records.UserNumbers), u.randomNumber)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	fmt.Println("Ваш рекорд сохранён.")

}
