package pancake

import (
	"BWINF/utils"
	"reflect"
	"testing"
)

func TestSortSteps_Push(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		p    SortSteps
		args args
		want SortSteps
	}{
		{
			name: "push step on empty",
			p:    SortSteps{},
			args: args{
				i: 1,
			},
			want: SortSteps{1},
		},
		{
			name: "push step on non empty",
			p:    SortSteps{1, 2, 3},
			args: args{
				i: 4,
			},
			want: SortSteps{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			p.Push(tt.args.i)

			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestSortSteps_Copy(t *testing.T) {
	tests := []struct {
		name string
		p    SortSteps
		want SortSteps
	}{
		{
			name: "testing copy",
			p:    SortSteps{1, 2, 3},
			want: SortSteps{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := *tt.p.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}

			got.Push(0)

			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("Copy() modified pancake, passed in %v, got %v", tt.p, tt.want)
			}
		})
	}
}

func TestSortSteps_String(t *testing.T) {
	type testCase[T utils.Number] struct {
		name string
		s    SortSteps
		want string
	}
	tests := []testCase[int]{
		{
			name: "testing empty",
			s:    SortSteps{},
			want: "",
		},
		{
			name: "testing one digit numbers",
			s:    SortSteps{1, 2, 3},
			want: "1 2 3",
		},
		{
			name: "testing two digit numbers",
			s:    SortSteps{10, 20, 30},
			want: "10 20 30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSortSteps(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[T utils.Number] struct {
		name string
		args args
		want SortSteps
	}
	tests := []testCase[int]{
		{
			name: "testing single number",
			args: args{
				s: "1",
			},
			want: SortSteps{1},
		},
		{
			name: "testing multiple numbers",
			args: args{
				s: "1 2 3",
			},
			want: SortSteps{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseSortSteps[int](tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSortSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
