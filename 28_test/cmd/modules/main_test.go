package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "#1",
			values: []int{1, 2},
			want:   3,
		},
		{
			name:   "#2",
			values: []int{1},
			want:   1,
		},
		{
			name:   "#3",
			values: []int{5, 5},
			want:   10,
		},
		{
			name:   "#3",
			values: []int{5, 5},
			want:   10,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if sum := Sum(v.values...); sum != v.want {
				t.Errorf("Sum() = %d, want %d", sum, v.want)
			}
		})
	}
}
