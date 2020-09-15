package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
  z := 1.0
  for i := 0; i < 10; i++ {
    fmt.Println(z)
    //To approximate the square root of a number, first you find the difference of what z^2 is from the goal
    //Then you divide by the derivative of z^2 to scale how much we change z by how much z^2 is changing (Newtons Method of numerical analysis)
    z -= (z*z - x) / (z*2)
	
  }
  return z
}

func main() {
	fmt.Println(Sqrt(20))
}