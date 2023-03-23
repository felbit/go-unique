package goniq_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/felbit/goniq"
)

// TODO: Use property based testing to test all possible types

func TestNewSetString(t *testing.T) {
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
			got := goniq.NewSet(tc.slice)
			assert.Equal(t, tc.expect, got)
		})
	}
}

func TestNewSetInt(t *testing.T) {
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
			got := goniq.NewSet(tc.slice)
			if !reflect.DeepEqual(tc.slice, tc.expect) {
				t.Errorf("NewSet(): got %v, expected %v", got, tc.expect)
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
			got := goniq.Add(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", got, tc.expect)
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
			got := goniq.Add(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", got, tc.expect)
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
			got := goniq.Add(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", got, tc.expect)
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
			got := goniq.Add(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Add(): got %v, expect %v", got, tc.expect)
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
			got := goniq.Remove(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", got, tc.expect)
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
			got := goniq.Remove(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", got, tc.expect)
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
			got := goniq.Remove(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", got, tc.expect)
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
			got := goniq.Remove(tc.slice, tc.element)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Remove(): got %v, expect %v", got, tc.expect)
			}
		})
	}
}
