package functional

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase[T any] struct {
	want T
	f    func() T
}

func (tc *testCase[T]) test(t *testing.T) {
	if diff := cmp.Diff(tc.want, tc.f()); diff != "" {
		t.Errorf("Unexpected result (-want, +got):\n%s", diff)
	}
}

type TestCase interface {
	test(*testing.T)
}

func TestFunctional(t *testing.T) {
	for _, test := range []struct {
		name string
		tc   TestCase
	}{
		{
			name: "if true",
			tc: &testCase[string]{
				"truth",
				func() string { return If(true, "truth", "lies") },
			},
		},
		{
			name: "if false",
			tc: &testCase[string]{
				"lies",
				func() string { return If(false, "truth", "lies") },
			},
		},
		{
			name: "any true",
			tc: &testCase[bool]{
				true,
				func() bool {
					return Any([]string{
						"hello",
						"there",
						"general",
						"kenobi",
					}, func(s string) bool {
						return strings.Contains(s, "here")
					})
				},
			},
		},
		{
			name: "any false",
			tc: &testCase[bool]{
				false,
				func() bool {
					return Any([]string{
						"hello",
						"there",
						"general",
						"kenobi",
					}, func(s string) bool {
						return strings.Contains(s, "z")
					})
				},
			},
		},
		{
			name: "Count",
			tc: &testCase[int]{
				2,
				func() int {
					return Count([]string{
						"un",
						"deux",
						"trois",
						"deux",
						"quatre",
					}, "deux")
				},
			},
		},
		{
			name: "CountFunc",
			tc: &testCase[int]{
				3,
				func() int {
					return CountFunc([]string{
						"un",
						"deux",
						"trois",
						"deux",
						"quatre",
					}, func(s string) bool {
						return strings.Contains(s, "e")
					})
				},
			},
		},
		{
			name: "Flat",
			tc: &testCase[[]int]{
				[]int{
					0,
					1,
					2,
					3,
					4,
					5,
					6,
				},
				func() []int {
					return Flat([][]int{
						{0, 1},
						{2},
						{3, 4, 5},
						{},
						{6},
						{},
					})
				},
			},
		},
		{
			name: "Filter",
			tc: &testCase[[]int]{
				[]int{
					0,
					2,
					6,
				},
				func() []int {
					return Filter([]int{0, 1, 2, 3, 4, 5, 6}, func(i int) bool {
						return i%2 == 0 && i != 4
					})
				},
			},
		},
		{
			name: "All returns true",
			tc: &testCase[bool]{
				true,
				func() bool {
					return All([]int{0, 1, 2, 3, 4, 5, 6}, func(i int) bool {
						return i <= 6
					})
				},
			},
		},
		{
			name: "All returns false",
			tc: &testCase[bool]{
				false,
				func() bool {
					return All([]int{0, 1, 2, 3, 4, 5, 6}, func(i int) bool {
						return i <= 6 && i != 3
					})
				},
			},
		},
		{
			name: "None returns true",
			tc: &testCase[bool]{
				true,
				func() bool {
					return None([]int{0, 1, 2, 3, 4, 5, 6}, func(i int) bool {
						return i > 6
					})
				},
			},
		},
		{
			name: "None returns false",
			tc: &testCase[bool]{
				false,
				func() bool {
					return None([]int{0, 1, 2, 3, 4, 5, 6}, func(i int) bool {
						return i == 2
					})
				},
			},
		},
		{
			name: "MapWithIndex",
			tc: &testCase[[]*indexType]{
				[]*indexType{
					{0, []string{"abc"}},
					{1, []string{"def"}},
					{2, []string{"ghij"}},
				},
				func() []*indexType {
					return MapWithIndex([]string{"abc", "def", "ghij"}, func(i int, v string) *indexType {
						return &indexType{
							i, []string{v},
						}
					})
				},
			},
		},
		{
			name: "Map",
			tc: &testCase[[]int]{
				[]int{3, 1, 4},
				func() []int {
					return Map([]string{"abc", "d", "efgh"}, func(v string) int {
						return len(v)
					})
				},
			},
		},
		{
			name: "Reduce",
			tc: &testCase[int]{
				720,
				func() int {
					return Reduce(1, []int{1, 2, 3, 4, 5, 6}, func(a, b int) int {
						return a * b
					})
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			test.tc.test(t)
		})
	}
}

type indexType struct {
	Idx int
	V   []string
}
