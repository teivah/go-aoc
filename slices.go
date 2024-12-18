package aoc

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func SliceCopy[T any](in []T) []T {
	res := make([]T, len(in))
	copy(res, in)
	return res
}

// DeleteSliceIndex deletes a given index.
func DeleteSliceIndex[T any](in []T, index int, deepCopy bool) []T {
	if deepCopy {
		in = SliceCopy(in)
	}
	return append(in[:index], in[index+1:]...)
}

// DeleteSliceIndices deletes a list of indices.
func DeleteSliceIndices[T any](in []T, indices []int, deepCopy bool) []T {
	if len(in) == 0 {
		return nil
	}
	if deepCopy {
		in = SliceCopy(in)
	}
	set := SliceToSet(indices)

	i := 0
	return slices.DeleteFunc(in, func(t T) bool {
		defer func() {
			i++
		}()
		return set[i]
	})
}

func IsSliceSorted[T constraints.Ordered](in []T, comp func(a T, b T) bool) bool {
	if len(in) <= 1 {
		return true
	}
	prev := in[0]
	for i := 1; i < len(in); i++ {
		cur := in[i]
		if !comp(prev, cur) {
			return false
		}
		prev = cur
	}
	return true
}

func IsSliceMonotonicallyIncreasing[T constraints.Ordered](in []T) bool {
	return IsSliceSorted(in, func(a T, b T) bool {
		return b >= a
	})
}

func IsSliceMonotonicallyDecreasing[T constraints.Ordered](in []T) bool {
	return IsSliceSorted(in, func(a T, b T) bool {
		return b <= a
	})
}

func IsSliceStrictlyIncreasing[T constraints.Ordered](in []T) bool {
	return IsSliceSorted(in, func(a T, b T) bool {
		return b > a
	})
}

func IsSliceStrictlyDecreasing[T constraints.Ordered](in []T) bool {
	return IsSliceSorted(in, func(a T, b T) bool {
		return b < a
	})
}

func SliceCount[T comparable](in []T) map[T]int {
	res := make(map[T]int)
	for _, v := range in {
		res[v]++
	}
	return res
}
