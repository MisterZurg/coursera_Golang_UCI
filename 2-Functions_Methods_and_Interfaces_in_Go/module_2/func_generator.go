/*
Let us assume the following formula for displacement s as a function of time t,
acceleration a, initial velocity vo, and initial displacement so.

s = ½ a t^2 + vo*t + so

Write a program which first prompts the user to enter values for acceleration,
initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial displacement so.
GenDisplaceFn() should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial displacement.
The function returned by GenDisplaceFn() should take one float64 argument t,
representing time, and return one float64 argument which is the displacement travelled after time t.
*/
package main

import (
	"fmt"
	"math"
)

// displacement s
// time - t, acceleration - a, initial velocity vo, and initial displacement so.
// s = 0,5 * a * t^2 * v0*t + so

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64
	fmt.Println("Enter values for acceleration,\n" +
		"initial velocity,\n" +
		"and initial displacement.")
	fmt.Scan(&acceleration, &initialVelocity, &initialDisplacement)
	fmt.Println("Enter time:")
	fmt.Scan(&time)

	Displacement := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(Displacement(time))
}

func GenDisplaceFn(acceleration, velocity, dsplacement float64) func(time float64) float64 {
	fn := func(insideTime float64) float64 {
		return 0.5*acceleration*math.Pow(insideTime, 2)*velocity*insideTime + dsplacement
	}
	return fn
}
