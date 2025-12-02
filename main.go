package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Name         string
	userNumber   int
	Try          int
	Limit        int
	randomNumber int
}

func main() {

	var name string
	fmt.Println("Введите ваше Имя:")
	fmt.Scanln(&name)

	user := User{
		Name: name,
	}
	user.startGame()
	fmt.Println("Твоя профиль:", user.Name)
	fmt.Println("Количество потраченных попыток:", user.Try)
}
func (u *User) genRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	u.randomNumber = rand.Intn(10) + 1
}

func (u *User) startGame() {
	u.Limit = 5
	u.Try = 0
	u.genRandomNumber()

	fmt.Println("Угадай число от 1 до 10")
	for u.Try < u.Limit {
		fmt.Println("Введите ваше число")
		fmt.Scanln(&u.userNumber)
		if u.userNumber < 1 || u.userNumber > 10 {
			fmt.Println("Ошибка ввода. Число должно быть от 1 до 10")
			continue
		}
		u.Try++
		if u.userNumber < u.randomNumber {
			fmt.Println("Бери повыше")
		} else if u.userNumber > u.randomNumber {
			fmt.Println("Маленько меньше")
		} else {
			fmt.Println("Красавчик ты угадал")
			return
		}
		if u.Try < u.Limit {
			fmt.Printf(" осталось %d попыток. \n", u.Limit-u.Try)
		} else {
			fmt.Printf("У тебя кончились попытки,загаданное число было %d\n", u.randomNumber)

		}
	}
}
