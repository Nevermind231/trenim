package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Record struct {
	Name        string
	Try         int
	UserNumbers []int
	Success     int
}
type Game struct {
	Guess        int
	randomNumber int
	Try          int
}

const MaxAttempts = 5
const RecordFile = "records.txt"

func main() {

	var records Record
	user := Game{
		Try: 0,
	}

	user.playGame(&records)

	if err := saveToFile(records, user); err != nil {
		fmt.Println("Ошибка сохранения:", err)
	} else {
		fmt.Println("Ваш рекорд сохранён.")
	}
}
func (u *Game) genRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	u.randomNumber = rand.Intn(10) + 1
}

func initGame(u *Game, records *Record) {
	u.genRandomNumber()
	u.Try = 0
	records.UserNumbers = []int{}
	records.Success = 0

}

func (u *Game) playGame(records *Record) {

	fmt.Println("Введите ваше Имя:")
	fmt.Scanln(&records.Name)

	initGame(u, records)

	fmt.Println("Угадай число от 1 до 10")
	for u.Try < MaxAttempts {
		guess, err := readNumber(1, 10)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		u.Guess = guess
		u.Try++
		records.UserNumbers = append(records.UserNumbers, u.Guess)

		if u.Try < MaxAttempts {
			fmt.Printf(" осталось %d попыток. \n", MaxAttempts-u.Try)
			if u.Guess < u.randomNumber {
				fmt.Println("Бери повыше")
			} else if u.Guess > u.randomNumber {
				fmt.Println("Маленько меньше")
			} else {
				fmt.Println("Красавчик ты угадал")
				records.Success = u.Try
				return
			}

		} else {
			fmt.Printf("У тебя кончились попытки,загаданное число было %d\n", u.randomNumber)

		}
	}
}
func readNumber(min, max int) (int, error) {
	var num int

	_, err := fmt.Scanln(&num)
	if err != nil {
		var diskard string
		fmt.Scanln(&diskard)
		return 0, fmt.Errorf("нужно ввести число")
	}

	if num < min || num > max {
		return 0, fmt.Errorf("число должно быть от %d до %d", min, max)
	}

	return num, nil
}

func saveToFile(records Record, u Game) error {
	file, err := os.OpenFile("records.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)

	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "Имя: %s Числа: ", records.Name)
	if err != nil {
		return fmt.Errorf("write name: %w", err)
	}

	for _, num := range records.UserNumbers {
		_, err := fmt.Fprintf(file, "%d ", num)
		if err != nil {
			return fmt.Errorf("write number: %w", err)
		}
	}

	if records.Success > 0 {
		_, err = fmt.Fprintf(file, "Угадал на попытке: %d Загаданное число: %d\n", records.Success, u.randomNumber)
	} else {
		_, err = fmt.Fprintf(file, "Не угадал число за %d попыток. Загаданное число: %d\n", len(records.UserNumbers), u.randomNumber)
		if err != nil {
			return fmt.Errorf("write result: %w", err)
		}
	}
	return nil
}
