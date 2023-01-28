package slice

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	type args[S any] struct {
		s []S
	}
	type testCase[S any] struct {
		name string
		args args[S]
		want []S
	}
	tests := []testCase[int]{
		{
			name: "reversing empty slice",
			args: args[int]{
				s: []int{},
			},
			want: []int{},
		},
		{
			name: "reversing slice with one element",
			args: args[int]{
				s: []int{1},
			},
			want: []int{1},
		},
		{
			name: "reversing slice with two elements",
			args: args[int]{
				s: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "reversing slice with three elements",
			args: args[int]{
				s: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSlice(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
