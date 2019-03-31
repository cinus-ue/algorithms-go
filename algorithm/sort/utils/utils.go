package utils

import (
	"math/rand"
	//"time"
)

func GetArrayOfSize(n int) []int {
	numbers := make([]int, 0)
	for i := 0; i < n; i++ {
		//numbers  = append(numbers , randInt())
		numbers = append(numbers, randInt(1, 1000))
	}
	return numbers[0:n]
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

//func randInt() int {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	return r.Intn(100)
//}

func Swap(array []int, index1 int, index2 int) {
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}
