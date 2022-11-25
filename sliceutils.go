package sliceutils

// Map applies the given function over every element in the slice and returns a new slice.
func Map[T any](arr []T, fn func(T) T) []T {
	ret := []T{}
	for _, i := range arr {
		ret = append(ret, fn(i))
	}
	return ret
}

// Filter applies the given function over every element in the slice and returns a new slice
// containing only the elements for which the function returned `true`.
func Filter[T any](arr []T, fn func(T) bool) []T {
	ret := []T{}
	for _, i := range arr {
		if fn(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

// Find returns the first element for which the given function returned `true`.
func Find[T any](arr []T, fn func(T) bool) (T, bool) {
	var ret T
	for _, i := range arr {
		if fn(i) {
			return i, true
		}
	}
	return ret, false
}

// Some applies the given function to every element in the slice and returns
// `true` if at least one of the invocations returned `true`.
func Some[T any](arr []T, fn func(T) bool) bool {
	for _, i := range arr {
		if fn(i) {
			return true
		}
	}
	return false
}

// Every applies the given function to every element in the slice and returns
// `true` if all the invocations returned `true`.
func Every[T any](arr []T, fn func(T) bool) bool {
	ret := true
	for _, i := range arr {
		if !fn(i) {
			ret = false
		}
	}
	return ret
}

// ForEach applies the given function to every element in the slice.
func ForEach[T any](arr []T, fn func(T)) {
	for _, i := range arr {
		fn(i)
	}
}

// Reduce applies the given function to every element in the slice, passing the return value of
// the previous invocation to the next one (in the first invocation, the _initial_ value is used).
func Reduce[T any, U any](arr []T, fn func(accumulator U, current T) U, initial U) U {
	ret := initial
	for _, i := range arr {
		ret = fn(ret, i)
	}
	return ret
}
