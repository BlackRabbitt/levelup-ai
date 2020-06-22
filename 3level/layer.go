/*
Level2: Dense Layer, Activation Fuction. In this section, following code represents Network with 2 layers (3 neurons and 1 neurons respectively).
This network will process inputs in batches.
*/
package main

import "math/rand"
import "time"
import "fmt"
import "log"
import . "github.com/BlackRabbitt/levelup-ai/shared"

type Layer interface {
	Forward(x [][]float32)
}

type DenseLayer struct {
	Weights [][]float32
	Biases []float32
	Output [][]float32
}

func NewDenseLayer(n_inputs int, n_neurons int) (denseLayer *DenseLayer) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	weights := make([][]float32, n_inputs)
	for i := range weights {
		weights[i] = make([]float32, n_neurons)
		for j := range weights[i] {
			weights[i][j] = 0.10 * r.Float32() // 0.10 is scale factor
		}
	}
	biases := make([]float32, n_neurons)
	for i := range biases {
		biases[i] = r.Float32()
	}

	denseLayer = &DenseLayer{Weights: weights, Biases: biases}

	return
}

// output = activation(dot(input, kernel) + bias)
func (denseLayer *DenseLayer) Forward(x [][]float32) {
	p, err := DotProduct(x, denseLayer.Weights)
	if err != nil {
		log.Fatal(err)
	}
	denseLayer.Output = Activation_ReLU(Add(p, denseLayer.Biases))
}

// step activation: (x >= 0 ? 1 : 0)
func Activation_Step(inputs [][]float32) (outputs [][]float32) {
	outputs = make([][]float32, len(inputs))
	for i, u := range inputs {
		outputs[i] = make([]float32, len(u))
		for j, v := range u {
			if v >= 0 {
				outputs[i][j] = 1
			} else {
				outputs[i][j] = 0
			}
		}
	}
	return
}

// ReLU activation: (x <= 0 ? 0 : x)
func Activation_ReLU(inputs [][]float32) (outputs [][]float32) {
	outputs = make([][]float32, len(inputs))
	for i, u := range inputs {
		outputs[i] = make([]float32, len(u))
		for j, v := range u {
			if v <= 0 {
				outputs[i][j] = 0
			} else {
				outputs[i][j] = v
			}
		}
	}
	return
}

func main() {
	// x is inputs in a batch of size 3.
	x := [][]float32{
		{1, 2, 3, 2.5},
		{2.0, 5.0, -1.0, 2.0},
		{-1.5, 2.7, 3.3, -0.8},
	}

	// this program represents network with two dense layer as follows
	layer1 := NewDenseLayer(4, 5)
	layer2 := NewDenseLayer(5, 2)

	// forward pass
	layer1.Forward(x)
	layer2.Forward(layer1.Output)

	fmt.Println(layer2.Output)
}
