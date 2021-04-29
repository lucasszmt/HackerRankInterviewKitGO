package main

import (
	"fmt"
	"github.com/lucasszmt/hackerRank/interviewkit/arrays"
)

func main() {
	t := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	fmt.Println(arrays.HourglassSum(t))
}
