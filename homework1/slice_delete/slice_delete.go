package slice_delete

import (
	"fmt"
)

func SliceDelete[T any](idx int, vals []T) []T {

	if idx < 0 || idx >= len(vals) {
		fmt.Printf("len:%d idx %d out of range! \n", len(vals), idx)
		return nil
	}

	tmp := vals[:]
	vals = append(vals[:idx], tmp[idx+1:]...)

	return vals
}
