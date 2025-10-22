package main

func main() {
    fmt.Println([]int{[1,0,0,0,1],2})
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    var count int
    for i := range flowerbed {
        if flowerbed[i] != 0 {
            continue
        }
        if i != 0 && flowerbed[i-1] != 0 {
            continue
        }
        if i != len(flowerbed)-1 && flowerbed[i+1] != 0 {
            continue
        }
        flowerbed[i] = 1
        count++
    }
    return count >= n
}