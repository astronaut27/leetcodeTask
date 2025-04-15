package main

import "fmt"

var letterMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var ans []string
	backtrack("", 0, digits, &ans)
	return ans
}

func backtrack(curr string, index int, digits string, ans *[]string) {
	// Базовый случай: если длина текущей строки совпадает с длиной digits
	if len(curr) == len(digits) {
		// Добавляем текущую комбинацию в результат
		*ans = append(*ans, curr)
		return
	}
	// Получаем буквы, соответствующие текущей цифре
	letters := letterMap[digits[index]]
	for i := 0; i < len(letters); i++ {
		// Добавляем текущую букву и переходим к следующей цифре
		backtrack(curr+string(letters[i]), index+1, digits, ans)
	}
}

func main() {
	// Пример вызова
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations(""))   // []
	fmt.Println(letterCombinations("2"))  // ["a","b","c"]
}
