package main

import (
	"fmt"
	"sort"
)

func main() {
	boxTypes := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	fmt.Println(maximumUnits(boxTypes, 10))

}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	unitCount := 0
	for _, boxType := range boxTypes {
		if boxType[0] < truckSize {
			truckSize -= boxType[0]
			unitCount += boxType[0] * boxType[1]
		} else {
			unitCount += truckSize * boxType[1]
			break
		}
	}
	return unitCount
}
