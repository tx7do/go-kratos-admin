package slice

import (
	"reflect"
	"sort"
	"testing"
)

func SortUint32(slice []uint32) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func TestMergeAndDeduplicate(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []uint32
		slice2   []uint32
		expected []uint32
	}{
		{
			name:     "No duplicates",
			slice1:   []uint32{1, 2, 3},
			slice2:   []uint32{4, 5, 6},
			expected: []uint32{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "With duplicates",
			slice1:   []uint32{1, 2, 3},
			slice2:   []uint32{3, 4, 5},
			expected: []uint32{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty slice1",
			slice1:   []uint32{},
			slice2:   []uint32{1, 2, 3},
			expected: []uint32{1, 2, 3},
		},
		{
			name:     "Empty slice2",
			slice1:   []uint32{1, 2, 3},
			slice2:   []uint32{},
			expected: []uint32{1, 2, 3},
		},
		{
			name:     "Nil slice1",
			slice1:   nil,
			slice2:   []uint32{1, 2, 3},
			expected: []uint32{1, 2, 3},
		},
		{
			name:     "Both slices empty",
			slice1:   []uint32{},
			slice2:   []uint32{},
			expected: []uint32{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeAndDeduplicate(tt.slice1, tt.slice2)
			SortUint32(result)
			SortUint32(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
