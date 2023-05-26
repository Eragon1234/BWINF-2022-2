package pancake

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestPancake_Flip(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		p    Stack
		args args
		want Stack
	}{
		{
			name: "flip only top most pancake",
			p:    Stack{1, 2, 3},
			args: args{
				i: 1,
			},
			want: Stack{1, 2},
		},
		{
			name: "flip all pancakes",
			p:    Stack{1, 2, 3},
			args: args{
				i: 3,
			},
			want: Stack{3, 2},
		},
		{
			name: "flip from the middle",
			p:    Stack{1, 2, 3},
			args: args{
				i: 2,
			},
			want: Stack{1, 3},
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
		i int8
	}
	tests := []struct {
		name string
		p    Stack
		args args
		want Stack
	}{
		{
			name: "push pancake on empty stack",
			p:    Stack{},
			args: args{
				i: 1,
			},
			want: Stack{1},
		},
		{
			name: "push pancake on non empty stack",
			p:    Stack{1, 2, 3},
			args: args{
				i: 4,
			},
			want: Stack{1, 2, 3, 4},
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

func TestParseStack(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Stack
		wantErr bool
	}{
		{
			name: "parse empty pancake",
			args: args{
				reader: strings.NewReader(""),
			},
			want:    Stack{},
			wantErr: false,
		},
		{
			name: "parse pancake with 1 element",
			args: args{
				reader: strings.NewReader("1\n1"),
			},
			want:    Stack{1},
			wantErr: false,
		},
		{
			name: "parse pancake with 2 elements",
			args: args{
				reader: strings.NewReader("2\n1\n2"),
			},
			want:    Stack{2, 1},
			wantErr: false,
		},
		{
			name: "parse pancake with 3 elements",
			args: args{
				reader: strings.NewReader("3\n1\n2\n3"),
			},
			want:    Stack{3, 2, 1},
			wantErr: false,
		},
		{
			name: "testing if invalid pancake returns an error",
			args: args{
				reader: strings.NewReader("this\nis\nan\ninvalid\npancake"),
			},
			want:    Stack{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStack(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPancake_Copy(t *testing.T) {
	tests := []struct {
		name string
		p    Stack
		want Stack
	}{
		{
			name: "testing copy",
			p:    Stack{1, 2, 3},
			want: Stack{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := *tt.p.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}

			got.Flip(0)

			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("Copy() modified pancake, passed in %v, got %v", tt.p, tt.want)
			}
		})
	}
}
