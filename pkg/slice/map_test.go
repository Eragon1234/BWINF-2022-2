package slice

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		s []T
		f func(T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[string, int]{
		{
			name: "testing empty slice",
			args: args[string, int]{
				s: []string{},
				f: func(s string) int { return len(s) },
			},
			want: []int{},
		},
		{
			name: "testing slice with one element",
			args: args[string, int]{
				s: []string{"a"},
				f: func(s string) int { return len(s) },
			},
			want: []int{1},
		},
		{
			name: "testing slice with two elements",
			args: args[string, int]{
				s: []string{"a", "abc"},
				f: func(s string) int { return len(s) },
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
