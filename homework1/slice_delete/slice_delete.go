package slice_delete

import (
	"fmt"
)

func SliceDelete[T any](idx int, vals []T) []T {

	if idx < 0 || idx >= len(vals) {
		fmt.Printf("len:%d idx %d out of range! \n", len(vals), idx)
		return nil
	}

	if len(vals)-1 <= cap(vals)/2 {
		newVals := make([]T, 0, cap(vals)/2)
		newVals = append(newVals, vals[:idx]...)
		newVals = append(newVals, vals[idx+1:]...)
		vals = newVals
	} else {
		vals = append(vals[:idx], vals[idx+1:]...)
	}

	return vals
}
