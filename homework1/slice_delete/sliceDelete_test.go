package slice_delete_test

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
// @date: 2023年10月3日
// @brief: 切片元素删除-缩容测试
// @author: Kewin Li
func TestSliceDelete(t *testing.T) {
	nums := make([]int, 20)

	t.Run("缩容测试", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			fmt.Printf("%d %d\n", len(nums), cap(nums))
			nums = sd.SliceDelete(0, nums)
		}
	})

}

// @func: TestSliceDelete
// @date: 2023年10月1日
// @brief: 切片元素删除-功能性测试
// @author: Kewin Li
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

	// 3. 删除中间元素触发缩容
	vals = res
	res = sd.SliceDelete[int](1, vals)
	fmt.Printf("%v %d %d\n", res, len(res), cap(res))

	// Output:
	// [2 3 4 5] 4 5
	// [2 3 4] 3 5
	// [2 4] 2 2

}
