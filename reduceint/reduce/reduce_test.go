package reduce

import "testing"

func TestReduceInt(t *testing.T) {
	type args struct {
		a []int
		f func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				f: func(x, y int) int { return x + y },
			},
			want: 15,
		},
		{
			name: "product",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				f: func(x, y int) int { return x * y },
			},
			want: 120,
		},
		{
			name: "max",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				f: func(x, y int) int {
					if x > y {
						return x
					}
					return y
				},
			},
			want: 5,
		},
		{
			name: "min",
			args: args{
				a: []int{5, 3, 4, 1, 2},
				f: func(x, y int) int {
					if x < y {
						return x
					}
					return y
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceInt(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("ReduceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
