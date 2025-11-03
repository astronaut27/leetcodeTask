package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5,1,5,0,-7}))
}

func largestAltitude(gain []int) int {
    var tmp,result int
    for i:=0;i<len(gain);i++{
        tmp+=gain[i]
        if result<tmp{
            result=tmp
        }
    }
    return result
}