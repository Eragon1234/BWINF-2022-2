package slice

import (
	"testing"
)

func TestIndexOfBiggestInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test first element biggest",
			args: args{
				s: []int{3, 2, 1},
			},
			want: 0,
		},
		{
			name: "Test last element biggest",
			args: args{
				s: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "Test middle element biggest",
			args: args{
				s: []int{1, 3, 2},
			},
			want: 1,
		},
		{
			name: "Test empty slice",
			args: args{
				s: []int{},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOfBiggest(tt.args.s); got != tt.want {
				t.Errorf("IndexOfBiggest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexOfBiggestNonSortedInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test first element biggest",
			args: args{
				s: []int{3, 2, 1},
			},
			want: -1,
		},
		{
			name: "Test last element biggest",
			args: args{
				s: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "Test biggest element is sorted should return index of second biggest",
			args: args{
				s: []int{3, 0, 2, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOfBiggestNonSortedNumber(tt.args.s); got != tt.want {
				t.Errorf("IndexOfBiggestNonSortedNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonSortedIndex(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test only first element sorted",
			args: args{
				s: []int{3, 1, 2},
			},
			want: 1,
		},
		{
			name: "Test first two elements sorted",
			args: args{
				s: []int{4, 3, 1, 2},
			},
			want: 2,
		},
		{
			name: "Test all elements sorted",
			args: args{
				s: []int{3, 2, 1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonSortedIndex(tt.args.s); got != tt.want {
				t.Errorf("NonSortedIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
