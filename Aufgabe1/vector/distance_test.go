package vector

import "testing"

func TestDistance(t *testing.T) {
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
			name: "test 1 unit",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: 0, Y: 1},
			},
			want: 1,
		},
		{
			name: "test 2 units",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: 0, Y: 2},
			},
			want: 2,
		},
		{
			name: "test 3th quadrant",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: -1, Y: -1},
			},
			want: 1.4142135623730951,
		},
		{
			name: "test 2th quadrant",
			args: args{
				from: Coordinate{X: 0, Y: 0},
				to:   Coordinate{X: -1, Y: 1},
			},
			want: 1.4142135623730951,
		},
		{
			name: "test non origin",
			args: args{
				from: Coordinate{X: 1, Y: 1},
				to:   Coordinate{X: 2, Y: 2},
			},
			want: 1.4142135623730951,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
