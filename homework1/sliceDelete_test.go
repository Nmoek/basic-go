package sliceDelete_test

import (
	"fmt"
	sd "homework1/slice_delete"
	"testing"
)

const MAX = 100000

var vals = make([]int, MAX)

func m_min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// @func: BenchmarkReapt
// @date: 2023年10月1日
// @brief: 切片元素删除-性能测试
// @author: Kewin Li
// @param: *testing.B b
func BenchmarkSliceDelete(b *testing.B) {

	for i := 0; i < b.N-1; i++ {
		res := sd.SliceDelete[int](0, vals)
		vals = res
	}
}

// @func: TestSliceDelete
// @date: 2023年10月1日
// @brief: 切片元素删除-功能性测试
// @author: Kewin Li
// @param: *testing.T t
func ExampleSliceDelete() {

	vals := []int{1, 2, 3, 4, 5}
	var res []int

	// 1. 删除首元素
	res = sd.SliceDelete[int](0, vals)

	fmt.Printf("%v %d %d\n", res, len(res), cap(res))

	// 2. 删除尾元素
	vals = res
	res = sd.SliceDelete[int](len(vals)-1, vals)
	fmt.Printf("%v %d %d\n", res, len(res), cap(res))

	vals = []int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		res = sd.SliceDelete[int](0, vals)
		vals = res
	}
	fmt.Printf("%v %d %d\n", res, len(res), cap(res))

	// Output:
	// [2 3 4 5] 4 5
	// [2 3 4] 3 5
	// [] 0 8

}
