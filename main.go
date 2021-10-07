package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getSlice() (numSlice []int, err error) {
	var input string
	var num int
	fmt.Println("Введите числа через пробел")
	input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF{
		err = errors.New("некорректный ввод")
		return
	}
	for _, v := range strings.Split(strings.TrimSpace(input), " ") {
		num, err = strconv.Atoi(v)
		if err != nil {
			err = errors.New("необходимо вводить только целые числа")
			return
		}
		numSlice = append(numSlice, num)
	}

	return
}


func sortSlice(inputSlice []int) []int {
	var tmp []int
	var min int
	for {
		flag := true
		for i := 1; i < len(inputSlice); i++ {
			if inputSlice[i] < inputSlice[i-1] {
				min = inputSlice[i]
				tmp = append(inputSlice[:i], inputSlice[i+1:]...)
				inputSlice = append([]int{min}, tmp...)
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return inputSlice
}



func main()  {
	slice, err := getSlice()
	if err != nil {
		fmt.Printf("ошибка ввода слайса: %s", err)
		os.Exit(1)
	}

	result := sortSlice(slice)
	if err != nil {
		fmt.Printf("ошибка сортировки слайса: %s", err)
	}
	fmt.Printf("результат: %v", result)

}
