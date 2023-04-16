package pkg

import "testing"

func TestMin(t *testing.T) {
	type args[T Number] struct {
		a T
		b T
	}
	type testCase[T Number] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "values are euqal",
			args: args[int]{a: 1, b: 1},
			want: 1,
		},
		{
			name: "a is smaller",
			args: args[int]{a: 1, b: 2},
			want: 1,
		},
		{
			name: "b is smaller",
			args: args[int]{a: 2, b: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
