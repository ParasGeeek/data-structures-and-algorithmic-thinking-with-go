// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"fmt"
	"math"
)

func findMax(A []int) (max int) {
	max = math.MinInt32
	for _, value := range A {
		if value > max {
			max = value
		}
	}
	return max
}

func findMin(A []int) (min int) {
	min = math.MaxInt32
	for _, value := range A {
		if value < min {
			min = value
		}
	}
	return min
}
func findMinAndMax(A []int) (min int, max int) {
	min, max = math.MaxInt32, math.MinInt32
	for _, value := range A {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	var A = []int{11, -4, 7, 8, -10}
	fmt.Println("Max: ", findMax(A))
	fmt.Println("Min: ", findMin(A))
	min, max := findMinAndMax(A)
	fmt.Println("Min: ", min, "Max: ", max)
}
