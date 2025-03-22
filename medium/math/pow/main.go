package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 3))
	fmt.Println(myPow(2.10000, 3))

}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := 1.0
	base := x
	for n > 0 {
		if n%2 == 1 { // Если степень нечетная, умножаем на текущий base
			result *= base
		}
		base *= base // Умножаем base на себя (x^2, x^4, x^8, ...)
		n /= 2       // Делим степень пополам
	}
	return result
}
