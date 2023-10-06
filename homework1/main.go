package main

import (
	"fmt"
	sd "homework1/slice_delete"
	"os"
)

// @func: functionalTest
// @date: 2023年10月1日
// @brief: 功能性测试
// @author: Kewin Li
func functionalTest() {

	vals := []int{1, 2, 3, 4, 5}
	var res []int

	res = sd.SliceDelete[int](0, vals)

	fmt.Printf("vals=%p %v %d %d, res1=%p %v %d %d\n", &vals, vals, len(vals), cap(vals), &res, res, len(res), cap(res))

	fmt.Printf("----------------\n")
	vals = res
	res = sd.SliceDelete[int](1, vals)

	fmt.Printf("vals=%p %v %d %d, res1=%p %v %d %d\n", &vals, vals, len(vals), cap(vals), &res, res, len(res), cap(res))

	fmt.Printf("----------------\n")

	vals = res
	res = sd.SliceDelete[int](len(vals)-1, vals)

	fmt.Printf("vals=%p %v %d %d, res1=%p %v %d %d\n", &vals, vals, len(vals), cap(vals), &res, res, len(res), cap(res))
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("input xxx 1 \n")
		return
	}

	switch args[1][0] {
	case '1':
		functionalTest()
	}

}
