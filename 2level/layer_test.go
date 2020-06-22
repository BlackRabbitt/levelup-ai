package main

import "testing"

func TestDotProduct(t *testing.T) {
	a := [][]float32{
		{2, 4, 1},
		{2, 3, 9},
	}

	b := [][]float32{
		{1, 2, 3},
		{3, 6, 1},
		{2, 4, 7},
	}

	c := [][]float32{
		{16, 32, 17},
		{29, 58, 72},
	}

	actual, err := dotProduct(a, b)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	for i := range actual {
		for j := range actual[i] {
			if actual[i][j] != c[i][j] {
				t.Errorf("Expected: %v, Got: %v", c[i][j], actual[i][j])
			}
		}
	}
}
