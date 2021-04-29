package arrays

import "math"

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
