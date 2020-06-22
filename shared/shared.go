package shared

import "fmt"

// add bias to each element
func Add(p [][]float32, biases []float32) (o [][]float32) {
	o = make([][]float32, len(p))
	for i, u := range p {
		o[i] = make([]float32, len(u))
		for j, v := range u {
			o[i][j] = v + biases[j]
		}
	}
	return
}

// matrix multiplication uses dotproduct to find the element in output matrix
func DotProduct(a [][]float32, b [][]float32) ([][]float32, error) {
	// dimension of a is (m x n)
	// dimension of b is (p x q)

	m := len(a)
	n := len(a[0])
	p := len(b)
	q := len(b[0])
	c := make([][]float32, m) // multiplication result


	if n != p {
		return c, fmt.Errorf("Shape Error: trying to multiply matrix with shape (%v x %v) and (%v x %v)", m, n, p, q)
	}

	for i :=0; i < m; i++ {
		c[i] = make([]float32, q)
		for j := 0; j < q; j ++ {
			var sum float32 = 0
			for k := 0; k < p; k ++ {
				sum += a[i][k] * b[k][j]
			}
			c[i][j] = sum
		}
	}
	return c, nil
}
