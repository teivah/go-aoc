package aoc

// InnerMapGet performs a get or init in an inner map.
func InnerMapGet[K1 comparable, K2 comparable, X any](m map[K1]map[K2]X, k K1, cap ...int) map[K2]X {
	v, ok := m[k]
	if !ok {
		if len(cap) != 0 {
			v = make(map[K2]X, cap[0])
		} else {
			v = make(map[K2]X)
		}
		m[k] = v
	}
	return v
}

func MapCopy[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
