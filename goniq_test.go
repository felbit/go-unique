package goniq_test

import (
	"reflect"
	"testing"

	"github.com/felbit/goniq"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		slice   []string
		element string
		want    []string
	}{
		{"add to empty slice", []string{}, "foo", []string{"foo"}},
		{"add to non-empty slice", []string{"bar", "baz"}, "foo", []string{"bar", "baz", "foo"}},
		{"add duplicate element", []string{"bar", "baz"}, "baz", []string{"bar", "baz"}},
		{"add element that already exists in slice", []string{"bar", "baz", "foo"}, "foo", []string{"bar", "baz", "foo"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := goniq.Add(tt.slice, tt.element)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		slice   []string
		element string
		want    []string
	}{
		{"remove from empty slice", []string{}, "foo", []string{}},
		{"remove from non-empty slice", []string{"bar", "baz", "foo"}, "baz", []string{"bar", "foo"}},
		{"remove non-existing element", []string{"bar", "baz"}, "foo", []string{"bar", "baz"}},
		{"remove last element", []string{"bar", "baz", "foo"}, "foo", []string{"bar", "baz"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := goniq.Remove(tt.slice, tt.element)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
