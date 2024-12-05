package aoc

import "strconv"

// StringToInt converts a string into an int.
// It panics in case of a parsing error.
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// TryStringToInt tries to convert a string into an int.
func TryStringToInt(s string) (int, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return i, true
}

// StringsToInts converts a slice of strings to a slice of ints.
// It panics in case of a parsing error.
func StringsToInts(s []string) []int {
	res := make([]int, len(s))
	for i, v := range s {
		res[i] = StringToInt(v)
	}
	return res
}

// RuneToInt converts a rune into an int.
func RuneToInt(r rune) int {
	return int(r - '0')
}

// IntToRune converts an int into a rune.
func IntToRune(i int) rune {
	return []rune(strconv.Itoa(i))[0]
}

// IsRuneDecimal checks whether a rune is a decimal number.
func IsRuneDecimal(r rune) bool {
	return r >= '0' && r <= '9'
}

// MapKeysToSlice converts the maps keys to a slice.
func MapKeysToSlice[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// MapValuesToSlice converts the map values to a slice.
func MapValuesToSlice[K comparable, V any](m map[K]V) []V {
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// SliceToMap converts a slice into a map.
func SliceToMap[K comparable, V any](s []K, f func(K) V) map[K]V {
	m := make(map[K]V, len(s))
	for _, v := range s {
		m[v] = f(v)
	}
	return m
}

// SliceToSet converts a slice into a set (a map where all the keys are the elements
// from the slice and the values are true).
func SliceToSet[T comparable](s []T) map[T]bool {
	m := make(map[T]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}
