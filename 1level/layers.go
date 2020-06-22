/*
1level layers
*/
package main

import "fmt"
import "math/rand"
import "time"

// represent single layer network with 3 neurons.
func singleLayer() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // random number generator

	inputs := [2]float32{0.3, 1.7} // two inputs. eg. [humidity_sensor, temparture_sensor]

	weights1 := [2]float32{r.Float32(), r.Float32()} // weights for first neuron
	weights2 := [2]float32{r.Float32(), r.Float32()} // weights for second neuron
	weights3 := [2]float32{r.Float32(), r.Float32()} // weights for third neuron

	var bias1 float32 = r.Float32() // bias for first neuron
	var bias2 float32 = r.Float32() // bias for first neuron
	var bias3 float32 = r.Float32() // bias for first neuron

	// output is the final results calculated by these neurons.
	// note: number of neurons = number of output. commonly this is represented as Shape.
	var outputs [3]float32
	outputs[0] = inputs[0] * weights1[0] + inputs[1] * weights1[1] + bias1 // first neuron in action
	outputs[1] = inputs[0] * weights2[0] + inputs[1] * weights2[1] + bias2 // second neuron in action
	outputs[2] = inputs[0] * weights3[0] + inputs[1] * weights3[1] + bias3 // third neuron in action

	fmt.Println(outputs)
}

func doubleLayer() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // random number generator

	inputs := [2]float32{0.3, 1.7} // two inputs. eg. [humidity_sensor, temparture_sensor]

	// first layer setup with 3 neurons. hidden layer
	weights1 := [2]float32{r.Float32(), r.Float32()} // weights for first neuron
	weights2 := [2]float32{r.Float32(), r.Float32()} // weights for second neuron
	weights3 := [2]float32{r.Float32(), r.Float32()} // weights for third neuron

	bias1 := r.Float32() // bias for first neuron
	bias2 := r.Float32() // bias for first neuron
	bias3 := r.Float32() // bias for first neuron

	// output is the final results calculated by these neurons.
	// note: number of neurons = number of output. commonly this is represented as Shape.
	var outputs [3]float32
	outputs[0] = inputs[0] * weights1[0] + inputs[1] * weights1[1] + bias1 // first neuron in action
	outputs[1] = inputs[0] * weights2[0] + inputs[1] * weights2[1] + bias2 // second neuron in action
	outputs[2] = inputs[0] * weights3[0] + inputs[1] * weights3[1] + bias3 // third neuron in action

	// second layer setup with 1 neurons. output layer
	weights := [2]float32{r.Float32(), r.Float32()}
	bias := r.Float32()

	// output is the signle value results calculated by the last hidden layer. This is the output of whole network.
	var output float32
	output = outputs[0] * weights[0] + outputs[1] * weights[1] + bias

	fmt.Println(output)
}

func dotProduct() {
	// inputs[0] * weights1[0] + inputs[1] * weights1[1] from ^ is dotProduct of inputs and weights.
	a := [3]int{1, 2, 3}
	b := [3]int{2, 3, 4}

	var dotProduct int
	dotProduct = a[0]*b[0] + a[1]*b[1] + a[2]*b[2]

	fmt.Println(dotProduct)
}

func main() {
	fmt.Println("single layer network: ")
	singleLayer()

	fmt.Println("double Layer network: ")
	doubleLayer()

	fmt.Println("dot Product")
	dotProduct()
}
