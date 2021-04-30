package arrays

import (
	"fmt"
	"math"
)

func HourglassSum(arr [][]int32) int32 {
	var max int32 = math.MinInt32
	var sum int32
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[i])-2; j++ {
			sum += arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}
	return max
}

func minimumBribes(q []int32) {
	var x = make([]int32, len(q))
	for i, _ := range q {
		x[i] = int32(i + 1)
	}
	bribeCount := 0

	for index, elem := range q {
		if elem != x[index] {
			if elem != x[index] && index < len(q)-2 {
				if x[index+2] == elem {
					bribeCount += 2
					x[index+2] = x[index+1]
					x[index+1] = x[index]
					x[index] = elem
				} else if x[index+1] == elem {
					bribeCount += 1
					x[index+1] = x[index]
					x[index] = elem
				} else {
					fmt.Println("Too chaotic")
					return
				}
			}
		}
		if index == len(q)-2 && elem != x[index] {
			bribeCount++
		}
	}
	fmt.Println(bribeCount)
}
