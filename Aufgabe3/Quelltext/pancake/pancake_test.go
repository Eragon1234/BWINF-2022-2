package pancake

import (
	"reflect"
	"testing"
)

func TestPancake_Flip(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		p    Pancake
		args args
		want Pancake
	}{
		{
			name: "flip only top most pancake",
			p:    Pancake{1, 2, 3},
			args: args{
				i: 1,
			},
			want: Pancake{1, 2},
		},
		{
			name: "flip all pancakes",
			p:    Pancake{1, 2, 3},
			args: args{
				i: 3,
			},
			want: Pancake{3, 2},
		},
		{
			name: "flip from the middle",
			p:    Pancake{1, 2, 3},
			args: args{
				i: 2,
			},
			want: Pancake{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			p.Flip(tt.args.i)

			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPancake_Push(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		p    Pancake
		args args
		want Pancake
	}{
		{
			name: "push pancake on empty stack",
			p:    Pancake{},
			args: args{
				i: 1,
			},
			want: Pancake{1},
		},
		{
			name: "push pancake on non empty stack",
			p:    Pancake{1, 2, 3},
			args: args{
				i: 4,
			},
			want: Pancake{1, 2, 3, 4},
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
