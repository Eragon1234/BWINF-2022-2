package vector

import (
	"reflect"
	"testing"
)

func TestCoordinate_String(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test 0, 0",
			fields: fields{
				X: 0,
				Y: 0,
			},
			want: "0.000000 0.000000",
		},
		{
			name: "test 1, 0",
			fields: fields{
				X: 1,
				Y: 0,
			},
			want: "1.000000 0.000000",
		},
		{
			name: "test 0.5, 0.5",
			fields: fields{
				X: 0.5,
				Y: 0.5,
			},
			want: "0.500000 0.500000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Coordinate{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCoordinate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Coordinate
		wantErr bool
	}{
		{
			name: "test 0, 0 without decimals",
			args: args{
				s: "0 0",
			},
			want: Coordinate{
				X: 0,
				Y: 0,
			},
		},
		{
			name: "test 0, 0 with decimals",
			args: args{
				s: "0.000000 0.000000",
			},
			want: Coordinate{
				X: 0,
				Y: 0,
			},
		},
		{
			name: "test invalid non number input",
			args: args{
				s: "a b",
			},
			want:    Coordinate{},
			wantErr: true,
		},
		{
			name: "test invalid delimiter",
			args: args{
				s: "0,0",
			},
			want:    Coordinate{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCoordinate(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCoordinate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCoordinate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
