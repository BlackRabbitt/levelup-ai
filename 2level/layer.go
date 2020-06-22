/*
Level2: Batches, Layers. In this section, following code represents Network with 2 layers
This network will process inputs in batches.
*/
package main

import "math/rand"
import "time"
import "fmt"
import "log"

type Layer interface {
	Forward(x [][]float32)
}

// MyLayer is an example of what layer looks like. In next level, we will discuss denseLayer => one of the popular layer used.
type MyLayer struct {
	Weights [][]float32
	Biases []float32
	Output [][]float32
}

func NewMyLayer(n_inputs int, n_neurons int) (myLayer *MyLayer) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	weights := make([][]float32, n_inputs)
	for i := range weights {
		weights[i] = make([]float32, n_neurons)
		for j := range weights[i] {
			weights[i][j] = r.Float32()
		}
	}
	biases := make([]float32, n_neurons)
	for i := range biases {
		biases[i] = r.Float32()
	}

	myLayer = &MyLayer{Weights: weights, Biases: biases}

	return
}

// output = dot(input, weights) + bias
func (myLayer *MyLayer) Forward(x [][]float32) {
	p, err := dotProduct(x, myLayer.Weights)
	if err != nil {
		log.Fatal(err)
	}
	myLayer.Output = add(p, myLayer.Biases)
}

// add bias to each element
func add(p [][]float32, biases []float32) (o [][]float32) {
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
func dotProduct(a [][]float32, b [][]float32) ([][]float32, error) {
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

func main() {
	// x is inputs in a batch of size 3.
	x := [][]float32{
		{1, 2, 3, 2.5},
		{2.0, 5.0, -1.0, 2.0},
		{-1.5, 2.7, 3.3, -0.8},
	}

	// this program represents network with two layers as follows
	layer1 := NewMyLayer(4, 5)
	layer2 := NewMyLayer(5, 2)

	// forward pass
	layer1.Forward(x)
	layer2.Forward(layer1.Output)

	fmt.Println(layer2.Output)
}
