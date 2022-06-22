package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] > len(arr) {
			result = append(result, arr[len(arr)-1])
			break
		}
		if arr[i] == arr[i+1] {
			copy(arr[i:], arr[i+1:])
			arr[len(arr)-1] = 0
			arr = arr[:len(arr)-1]
		}
		if arr[i] == arr[i+1] {
			i--
		} else if arr[i] != arr[i+1] {
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)
}

// Прибрати всі дублікати з слайсу int.
// Приклад даних на вхід: [3, 4, 4, 3, 6, 3]
// видаляємо 3 по індексу 0
// видаляємо 4 по індексу 1
// видаляємо 3 по індексу 3
// Правильний результат: [3, 4, 6]
// Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
// func main(){
//     arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
//     var result []int
//     //...
//     // тут має бути ваш код
//     // змінна result в кінці функції має тримати слайс з вже видаленими дублікатами відповідно до правил
//
