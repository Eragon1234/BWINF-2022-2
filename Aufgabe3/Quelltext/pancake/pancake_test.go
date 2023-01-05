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

func TestParsePancakeFromReader(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Pancake
		wantErr bool
	}{
		{
			name: "parse empty pancake",
			args: args{
				reader: strings.NewReader(""),
			},
			want:    Pancake{},
			wantErr: false,
		},
		{
			name: "parse pancake with 1 element",
			args: args{
				reader: strings.NewReader("1\n1"),
			},
			want:    Pancake{1},
			wantErr: false,
		},
		{
			name: "parse pancake with 2 elements",
			args: args{
				reader: strings.NewReader("2\n1\n2"),
			},
			want:    Pancake{2, 1},
			wantErr: false,
		},
		{
			name: "parse pancake with 3 elements",
			args: args{
				reader: strings.NewReader("3\n1\n2\n3"),
			},
			want:    Pancake{3, 2, 1},
			wantErr: false,
		},
		{
			name: "testing if invalid pancake returns an error",
			args: args{
				reader: strings.NewReader("this\nis\nan\ninvalid\npancake"),
			},
			want:    Pancake{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePancakeFromReader(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePancakeFromReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePancakeFromReader() got = %v, want %v", got, tt.want)
			}
		})
	}
}
