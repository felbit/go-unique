package goniq_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/felbit/goniq"
)

// TODO: Use property based testing to test all possible types

func TestContains(t *testing.T) {
	type args[T goniq.Ordered] struct {
		s *[]T
		e T
	}
	type testCase[T goniq.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{"empty slice", args[int]{&[]int{}, 1}, false},
		{"non-empty slice", args[int]{&[]int{1, 2, 3}, 2}, true},
		{"non-existing element", args[int]{&[]int{1, 2}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, goniq.Contains(tt.args.s, tt.args.e), "Contains(%v, %v)", tt.args.s, tt.args.e)
		})
	}
}

func TestRemoveDuplicatesString(t *testing.T) {
	testCases := []struct {
		title  string
		slice  []string
		expect []string
	}{
		{"empty slice", []string{}, []string{}},
		{"1 elem slice", []string{"foo"}, []string{"foo"}},
		{"2 distinct elem slice", []string{"foo", "bar"}, []string{"foo", "bar"}},
		{"multiple duplicates", []string{"foo", "bar", "foo", "bar", "quuz"}, []string{"foo", "bar", "quuz"}},
		{"preserve original order", []string{"quuz", "foo", "quuz", "bar", "foo"}, []string{"quuz", "foo", "bar"}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.RemoveDuplicates(&tc.slice)
			assert.Equal(t, tc.expect, tc.slice)
		})
	}
}

func TestRemoveDuplicatesInt(t *testing.T) {
	testCases := []struct {
		title  string
		slice  []int
		expect []int
	}{
		{"empty slice", []int{}, []int{}},
		{"1 elem slice", []int{5}, []int{5}},
		{"2 distinct elem slice", []int{2, 5}, []int{2, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.RemoveDuplicates(&tc.slice)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("RemoveDuplicates(): got %v, expected %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestAddString(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []string
		element string
		expect  []string
	}{
		{"add to empty slice", []string{}, "foo", []string{"foo"}},
		{"add to non-empty slice", []string{"bar", "baz"}, "foo", []string{"bar", "baz", "foo"}},
		{"add duplicate element", []string{"bar", "baz"}, "baz", []string{"bar", "baz"}},
		{"add element that already exists in slice", []string{"bar", "baz", "foo"}, "foo", []string{"bar", "baz", "foo"}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Add(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestAddInt(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []int
		element int
		expect  []int
	}{
		{"add to empty slice", []int{}, 1, []int{1}},
		{"add to non-empty slice", []int{1, 2}, 3, []int{1, 2, 3}},
		{"add duplicate element", []int{1, 2}, 2, []int{1, 2}},
		{"add element that already exists in slice", []int{2, 3, 1}, 2, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Add(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestAddUint(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []uint
		element uint
		expect  []uint
	}{
		{"add to empty slice", []uint{}, 1, []uint{1}},
		{"add to non-empty slice", []uint{1, 2}, 3, []uint{1, 2, 3}},
		{"add duplicate element", []uint{1, 2}, 2, []uint{1, 2}},
		{"add element that already exists in slice", []uint{2, 3, 1}, 2, []uint{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Add(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestAddFloat(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []float32
		element float32
		expect  []float32
	}{
		{"add to empty slice", []float32{}, 1.3, []float32{1.3}},
		{"add to non-empty slice", []float32{1.2, 2.3}, 3.4, []float32{1.2, 2.3, 3.4}},
		{"add duplicate element", []float32{1.2, 2.3}, 2.3, []float32{1.2, 2.3}},
		{"add element that already exists in slice", []float32{2.4, 2.3, 2.1}, 2.3, []float32{2.1, 2.3, 2.4}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Add(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestAppendString(t *testing.T) {
	testCases := []struct {
		title    string
		slice    []string
		elements []string
		expect   []string
	}{
		{"add to empty slice", []string{}, []string{"foo", "bar"}, []string{"foo", "bar"}},
		{"add to non-empty slice", []string{"bar", "baz"}, []string{"foo", "bar"}, []string{"bar", "baz", "foo"}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Append(&tc.slice, tc.elements...)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestAppendInt(t *testing.T) {
	testCases := []struct {
		title    string
		slice    []int
		elements []int
		expect   []int
	}{
		{"add to empty slice", []int{}, []int{1, 2}, []int{1, 2}},
		{"add to non-empty slice", []int{1, 2}, []int{3, 2}, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Append(&tc.slice, tc.elements...)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []string
		element string
		expect  []string
	}{
		{"remove from empty slice", []string{}, "foo", []string{}},
		{"remove from non-empty slice", []string{"bar", "baz", "foo"}, "baz", []string{"bar", "foo"}},
		{"remove non-existing element", []string{"bar", "baz"}, "foo", []string{"bar", "baz"}},
		{"remove last element", []string{"bar", "baz", "foo"}, "foo", []string{"bar", "baz"}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Remove(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestRemoveInt(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []int
		element int
		expect  []int
	}{
		{"remove from empty slice", []int{}, 1, []int{}},
		{"remove from non-empty slice", []int{1, 2, 3}, 2, []int{1, 3}},
		{"remove non-existing element", []int{1, 2}, 3, []int{1, 2}},
		{"remove last element", []int{1, 2, 3}, 3, []int{1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Remove(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestRemoveUint(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []uint
		element uint
		expect  []uint
	}{
		{"remove from empty slice", []uint{}, 1, []uint{}},
		{"remove from non-empty slice", []uint{1, 2, 3}, 2, []uint{1, 3}},
		{"remove non-existing element", []uint{1, 2}, 3, []uint{1, 2}},
		{"remove last element", []uint{1, 2, 3}, 3, []uint{1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Remove(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestRemoveFloat(t *testing.T) {
	testCases := []struct {
		title   string
		slice   []float32
		element float32
		expect  []float32
	}{
		{"remove from empty slice", []float32{}, 0.1, []float32{}},
		{"remove from non-empty slice", []float32{0.1, 0.2, -3}, 0.2, []float32{-3, 0.1}},
		{"remove non-existing element", []float32{0.1, 0.2}, 0.19, []float32{0.1, 0.2}},
		{"remove last element", []float32{1.1, 2.2, 3.3}, 3.3, []float32{1.1, 2.2}},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.Remove(&tc.slice, tc.element)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", tc.slice, tc.expect)
			}
		})
	}
}

func TestRemoveStringsAkin(t *testing.T) {
	testCases := []struct {
		title    string
		slice    []string
		str      string
		expected []string
	}{
		{
			title:    "no matches",
			slice:    []string{"apple", "banana", "cherry"},
			str:      "dog",
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			title:    "one match",
			slice:    []string{"apple", "banana", "cherry"},
			str:      "b",
			expected: []string{"apple", "cherry"},
		},
		{
			title:    "multiple matches",
			slice:    []string{"apple", "banana", "cherry"},
			str:      "a",
			expected: []string{"cherry"},
		},
		{
			title:    "all matches",
			slice:    []string{"apple", "banana", "cherry"},
			str:      "a|e",
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			goniq.RemoveStringsAkin(&tc.slice, tc.str)
			if !reflect.DeepEqual(tc.slice, tc.expected) {
				t.Errorf("got %v, want %v", tc.slice, tc.expected)
			}
		})
	}
}
