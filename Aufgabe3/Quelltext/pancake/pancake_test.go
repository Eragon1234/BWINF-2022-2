package pancake

import (
	"reflect"
	"testing"
)

func TestPancake_Flip(t *testing.T) {
	type fields struct {
		stack []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "flip only top most pancake",
			fields: fields{
				stack: []int{1, 2, 3},
			},
			args: args{
				i: 1,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "flip all pancakes",
			fields: fields{
				stack: []int{1, 2, 3},
			},
			args: args{
				i: 3,
			},
			want: []int{3, 2, 1},
		},
		{
			name: "flip from the middle",
			fields: fields{
				stack: []int{1, 2, 3},
			},
			args: args{
				i: 2,
			},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pancake{
				Stack: tt.fields.stack,
			}
			p.Flip(tt.args.i)

			if !reflect.DeepEqual(p.Stack, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", p.Stack, tt.want)
			}
		})
	}
}

func TestPancake_Push(t *testing.T) {
	type fields struct {
		stack []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "push pancake on empty stack",
			fields: fields{
				stack: []int{},
			},
			args: args{
				i: 1,
			},
			want: []int{1},
		},
		{
			name: "push pancake on non empty stack",
			fields: fields{
				stack: []int{1, 2, 3},
			},
			args: args{
				i: 4,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pancake{
				Stack: tt.fields.stack,
			}
			p.Push(tt.args.i)

			if !reflect.DeepEqual(p.Stack, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", p.Stack, tt.want)
			}
		})
	}
}
