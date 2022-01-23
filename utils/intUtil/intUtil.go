package intUtil

import (
	"log"
	"strconv"
)

func Contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Join(nums []int) int {
	var str string
	for i := range nums {
		str += strconv.Itoa(nums[i])
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Println("int join err: ", err)
	}
	return num
}
