package utils

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	type testCase[T any] struct {
		name  string
		args  args[T]
		want  T
		want1 []T
	}
	tests := []testCase[int]{
		{
			name: "test pop with 3 elements",
			args: args[int]{
				s: []int{1, 2, 3},
			},
			want:  3,
			want1: []int{1, 2},
		},
		{
			name: "test pop with 2 elements",
			args: args[int]{
				s: []int{1, 2},
			},
			want:  2,
			want1: []int{1},
		},
		{
			name: "test pop with 1 element",
			args: args[int]{
				s: []int{1},
			},
			want:  1,
			want1: []int{},
		},
		{
			name: "test pop on empty slice",
			args: args[int]{
				s: []int{},
			},
			want:  0,
			want1: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pop(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
