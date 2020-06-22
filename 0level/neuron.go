/*
0level Neuron. In this section, following program represents single neuron that takes 2 inputs and outputs single.
*/
package main

import "fmt"
import "math/rand"
import "time"

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // random number generator

	inputs := [2]float32{0.3, 1.7} // two inputs. eg. [humidity_sensor, temparture_sensor]
	weights := [2]float32{r.Float32(), r.Float32()} // assign random weights to this each input
	var bias float32 = 2 // add bias

	// output is the final result/calculation done by a single neuron.
	var output float32
	output = inputs[0] * weights[0] + inputs[1] * weights[1] + bias

	fmt.Println(output)
}
