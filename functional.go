// functional implements functional programming utilities.
package functional

import "golang.org/x/exp/slices"

// If returns t if v is true; f otherwise
func If[T any](v bool, t, f T) T {
	if v {
		return t
	}
	return f
}

// Any returns true if f(t) == true for any t in ts.
func Any[T any](ts []T, f func(t T) bool) bool {
	for _, t := range ts {
		if f(t) {
			return true
		}
	}
	return false
}

func Count[T comparable](ts []T, k T) int {
	return CountFunc(ts, func(t T) bool { return t == k })
}

func CountFunc[T any](ts []T, f func(T) bool) int {
	return Reduce(0, ts, func(b int, t T) int {
		if f(t) {
			return b + 1
		}
		return b
	})
}

func Count2D[T comparable](tss [][]T, k T) int {
	return Count2DFunc(tss, func(t T) bool { return t == k })
}

func Count2DFunc[T any](tss [][]T, f func(T) bool) int {
	return Reduce(0, tss, func(b int, ts []T) int {
		return b + CountFunc(ts, f)
	})
}

func Flat[T comparable](tts [][]T) []T {
	return Reduce(nil, tts, func(a, b []T) []T {
		return append(a, b...)
	})
}

func Filter[T any](ts []T, f func(t T) bool) []T {
	var r []T
	for _, t := range ts {
		if f(t) {
			r = append(r, t)
		}
	}
	return r
}

// All Returns true if f(t) == true for all t in ts.
func All[T any](ts []T, f func(t T) bool) bool {
	for _, t := range ts {
		if !f(t) {
			return false
		}
	}
	return true
}

// None returns true if f(t) == false for all t in ts.
func None[T any](ts []T, f func(t T) bool) bool {
	for _, t := range ts {
		if f(t) {
			return false
		}
	}
	return true
}

// MapWithIndex maps all of the elements in `items` with function f.
func MapWithIndex[I, O any](items []I, f func(int, I) O) []O {
	var r []O
	for idx, item := range items {
		r = append(r, f(idx, item))
	}
	return r
}

// Map maps all of the elements in `items` with the function f.
func Map[I, O any](items []I, f func(I) O) []O {
	return MapWithIndex(items, func(idx int, i I) O {
		return f(i)
	})
}

// Reduce reduces `base` across `items` with function f.
func Reduce[B, T any](base B, items []T, f func(B, T) B) B {
	b := base
	for _, t := range items {
		b = f(b, t)
	}
	return b
}

func SortFunc[T any](items []T, lessFunc func(T, T) bool) {
	slices.SortFunc(items, func(a, b T) int {
		if lessFunc(a, b) {
			return -1
		}
		return 1
	})
}
