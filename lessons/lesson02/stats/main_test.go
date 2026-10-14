package main

import "testing"

func TestStats(t *testing.T) {
	tests := []struct {
		in  []int
		lo  int
		hi  int
		sum int
		avg float64
		ok  bool
	}{
		{[]int{12000, 7500, 31000, 4500}, 4500, 31000, 55000, 13750, true},
		{[]int{500}, 500, 500, 500, 500, true},
		{[]int{3, 1, 2}, 1, 3, 6, 2, true},
		{[]int{}, 0, 0, 0, 0, false},
		{nil, 0, 0, 0, 0, false},
	}
	for _, tt := range tests {
		lo, hi, sum, avg, ok := Stats(tt.in)
		if lo != tt.lo || hi != tt.hi || sum != tt.sum || avg != tt.avg || ok != tt.ok {
			t.Errorf("Stats(%v) = %d, %d, %d, %v, %v; want %d, %d, %d, %v, %v",
				tt.in, lo, hi, sum, avg, ok, tt.lo, tt.hi, tt.sum, tt.avg, tt.ok)
		}
	}
}
