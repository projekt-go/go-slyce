package slyce

func Map[T any, Y any](slice []T, fn func(elem T) Y) []Y {
	res := make([]Y, len(slice))
	for i, v := range slice {
		res[i] = fn(v)
	}
	return res
}

// doesn't preserve the order of the original slice
func RemoveUnordered[T any](slice []T, i int) []T {
	n := len(slice)
	if i != n-1 {
		slice[i] = slice[n-1]
	}
	return slice[:n-1]
}

// zero-allocation fast `Filter` but doesn't preserve order
// of the original slice, and it mutates the slice itself.
func Filter[T any](slice []T, predicate func(elem T) bool) []T {
	for i := 0; i < len(slice); i++ {
		if !predicate(slice[i]) {
			slice = RemoveUnordered(slice, i)
			i--
		}
	}
	return slice
}
