package check

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		size  int
		want  [][]int
	}{
		{
			name:  "chunk of size 2",
			slice: []int{1, 2, 3, 4, 5},
			size:  2,
			want:  [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:  "chunk of size 3",
			slice: []int{1, 2, 3, 4, 5},
			size:  3,
			want:  [][]int{{1, 2, 3}, {4, 5}},
		},
		{
			name:  "chunk of size 1",
			slice: []int{1, 2, 3},
			size:  1,
			want:  [][]int{{1}, {2}, {3}},
		},
		{
			name:  "chunk of size larger than slice",
			slice: []int{1, 2, 3},
			size:  5,
			want:  [][]int{{1, 2, 3}},
		},
		{
			name:  "chunk of size zero",
			slice: []int{1, 2, 3},
			size:  0,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.slice, tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
