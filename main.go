package main

import (
	"errors"
	"fmt"
	"os"
)

func getSlice()  (res []int, err error) {
	var size, num int
	fmt.Println("Введите целое положительное число - длинну списка")


	_, err = fmt.Scan(&size)

	if err != nil || size < 1 {
		err = errors.New("длина массива должна быть больше 0")
		return
	}
	fmt.Println("С новой строки вводите целые числа")
	for i := 0; i < size; i++ {
		_, err = fmt.Scan(&num)
		if err != nil {
			err = errors.New("элементом массива должно быть целое число")
			return
		}
		res = append(res, num)
	}
	return
}

func sortSlice(userSlice []int) (sortedSlice []int) {
	var minPosition int
	for {
		for i := range userSlice {
			if userSlice[minPosition] > userSlice[i] {
				minPosition = i
			}
		}
		sortedSlice = append(sortedSlice, userSlice[minPosition])
		userSlice = append(userSlice[:minPosition], userSlice[minPosition+1:]...)
		if len(userSlice) == 0 {
			break
		} else {
			minPosition = 0
		}
	}
	return
}

func main()  {
	userSlice, err := getSlice()
	if err != nil {
		fmt.Printf("ошибка ввода массива, %s", err)
		os.Exit(1)
	}
	fmt.Printf("Отсортированный срез: %v", sortSlice(userSlice))
}