package slice

import (
	"reflect"
	"testing"
)

func TestMakeFunc(t *testing.T) {
	type args[T any] struct {
		len int
		f   func(i int) T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test length 0",
			args: args[int]{
				len: 0,
				f: func(i int) int {
					return i
				},
			},
			want: []int{},
		},
		{
			name: "test func returns i",
			args: args[int]{
				len: 5,
				f: func(i int) int {
					return i
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "test func returns i+1",
			args: args[int]{
				len: 5,
				f: func(i int) int {
					return i + 1
				},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeFunc(tt.args.len, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
