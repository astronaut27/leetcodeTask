package main

import "fmt"

// Шаг 2: Вспомогательная функция для НОД (Алгоритм Евклида)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
    // 1. Проверка существования: str1 + str2 == str2 + str1
	if str1 + str2 != str2 + str1 {
		return ""
	}

	// 2. Находим НОД длин
	len1 := len(str1)
	len2 := len(str2)
	lx := gcd(len1, len2)

	// 3. Возвращаем префикс str1 длиной Lx
	return str1[:lx]
}

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // Вывод: AB
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))   // Вывод: ABC
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // Вывод: ""
}