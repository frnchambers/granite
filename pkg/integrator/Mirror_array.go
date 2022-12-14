package integrator

func make_mirror_array[T any](array []T, is_even bool) []T {
	rev := make_reverse_array(array)
	if is_even {
		return append(array, rev...)
	} else {
		return append(array, rev[1:]...)
	}
}

func make_reverse_array[T any](array []T) (reverse_array []T) {
	n := len(array)
	reverse_array = make([]T, n)
	for i := 0; i < n; i++ {
		reverse_array[i] = array[n-1-i]
	}
	return
}
