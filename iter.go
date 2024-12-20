package aoc

import "iter"

func Over[T any](ts ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range ts {
			if !yield(t) {
				return
			}
		}
	}
}
