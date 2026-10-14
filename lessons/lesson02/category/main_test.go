package main

import "testing"

func TestCategory(t *testing.T) {
	tests := []struct {
		amount  int
		want    string
		wantErr bool
	}{
		{-1, "", true},
		{0, "small", false},
		{9_999, "small", false},
		{10_000, "regular", false},
		{49_999, "regular", false},
		{50_000, "large", false},
	}
	for _, tt := range tests {
		got, err := Category(tt.amount)
		if got != tt.want || (err != nil) != tt.wantErr {
			t.Errorf("Category(%d) = %q, %v; want %q, wantErr %v", tt.amount, got, err, tt.want, tt.wantErr)
		}
	}
}
