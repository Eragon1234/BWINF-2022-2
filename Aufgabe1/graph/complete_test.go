package graph

import (
	"BWINF/Aufgabe1/vector"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParseComplete(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    WeightedGraph[vector.Coordinate, DistanceAngle]
		wantErr bool
	}{
		{
			name: "test 2 vertices",
			args: args{
				reader: strings.NewReader("2 2\n0 0"),
			},
			want: WeightedGraph[vector.Coordinate, DistanceAngle]{
				Vertices: map[string]Vertex[vector.Coordinate]{
					"0.000000 0.000000": {
						Name:  "0.000000 0.000000",
						Index: 1,
						Value: vector.Coordinate{X: 0, Y: 0},
					},
					"2.000000 2.000000": {
						Name:  "2.000000 2.000000",
						Index: 0,
						Value: vector.Coordinate{X: 2, Y: 2},
					},
				},
				adjacencyMatrix: [][]Edge[DistanceAngle]{
					{
						{
							From: 0,
							To:   0,
						},
						{
							From: 0,
							To:   1,
							Weight: DistanceAngle{
								Distance: 2.8284271247461903,
								Angle:    45,
							},
							Exists: true,
						},
					},
					{
						{
							From: 1,
							To:   0,
							Weight: DistanceAngle{
								Distance: 2.8284271247461903,
								Angle:    45,
							},
							Exists: true,
						},
						{
							From: 1,
							To:   1,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseComplete(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseComplete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseComplete() got = %v, want %v", got, tt.want)
			}
		})
	}
}
