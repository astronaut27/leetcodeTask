package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(asteroidsDestroyed(10, []int{3, 9, 19, 5, 21})) //true
	fmt.Println(asteroidsDestroyed(5, []int{4, 9, 23, 4}))      //false

}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	currMass := mass
	for _, a := range asteroids {
		if a > currMass {
			return false
		}
		currMass += a
	}
	return true
}
