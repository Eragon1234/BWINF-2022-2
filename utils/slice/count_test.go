package slice

import "testing"

func TestCount(t *testing.T) {
	type args[T comparable] struct {
		s   []T
		val T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "test empty slice",
			args: args[int]{s: []int{}, val: 1},
			want: 0,
		},
		{
			name: "test slice with one element",
			args: args[int]{s: []int{1}, val: 1},
			want: 1,
		},
		{
			name: "test zero occurrences",
			args: args[int]{s: []int{1, 2, 3, 4, 5}, val: 0},
		},
		{
			name: "test multiple occurrences",
			args: args[int]{s: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, val: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s, tt.args.val); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountFunc(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "test empty slice",
			args: args[int]{s: []int{}, f: func(i int) bool { return i == 0 }},
			want: 0,
		},
		{
			name: "test slice with one element",
			args: args[int]{s: []int{1}, f: func(i int) bool { return i == 1 }},
			want: 1,
		},
		{
			name: "test zero occurrences",
			args: args[int]{s: []int{1, 2, 3, 4, 5}, f: func(i int) bool { return i == 0 }},
			want: 0,
		},
		{
			name: "test multiple occurrences",
			args: args[int]{s: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, f: func(i int) bool { return i == 1 }},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFunc(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("CountFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
