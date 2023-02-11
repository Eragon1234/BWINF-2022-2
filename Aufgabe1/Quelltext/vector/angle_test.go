package vector

import "testing"

func TestAngle(t *testing.T) {
	type args struct {
		from Coordinate
		to   Coordinate
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 45 degrees",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: 1, Y: 1},
			},
			want: 45,
		},
		{
			name: "test 90 degrees",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: 1, Y: 0},
			},
			want: 90,
		},
		{
			name: "test 135 degrees",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: 1, Y: -1},
			},
			want: 135,
		},
		{
			name: "test 180 degrees",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: 0, Y: -1},
			},
			want: 180,
		},
		{
			name: "test 225 degrees",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: -1, Y: -1},
			},
			want: -135,
		},
		{
			name: "test non origin",
			args: args{
				from: Coordinate{X: 1, Y: 1},
				to:   Coordinate{X: 2, Y: 2},
			},
			want: 45,
		},
		{
			name: "test other direction",
			args: args{
				from: Coordinate{X: 2, Y: 2},
				to:   Coordinate{X: 0, Y: 0},
			},
			want: -135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Angle(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}
