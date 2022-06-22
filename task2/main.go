package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string

	stringToArr := strings.Split(input, " ")
	var arrToInt []int
	for _, v := range stringToArr {
		num, err := strconv.Atoi(v)
		if err == nil {
			arrToInt = append(arrToInt, num)
		}
	}

	sort.Ints(arrToInt)

	maxValue := arrToInt[len(arrToInt)-1]
	minValue := arrToInt[0]

	if len(arrToInt) == 1 {
		result = strconv.Itoa(minValue)
	} else {
		result = strconv.Itoa(maxValue) + " " + strconv.Itoa(minValue)
	}
	fmt.Println(result)
}

// На вхід подано стрінг з цілими числа, котри розділені пробілами.
// Треба повернути найбільше та найменше число.
// Наприклад:
// input := "1 2 3 4 5" // повертає "5 1"
// input := "1 9 3 4 -5" // повертає "9 -5"
// Уточнення:
// 1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
// 2. В стрінгі завжди буде принаймні одне число.
// 3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
// число). Найбільше число має бути першим.
// 4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
// func main(){
//     input := "1 9 3 4 -5"
//     var result string
//     //...
//     // тут має бути ваш код
//     // змінна result в кінці функції має тримати стрінг з правильними результатами, згідно до умови задачі
//
//
