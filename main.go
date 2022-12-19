package main

import "fmt"

func main(){
	var a float64
		fmt.Print("Please enter acceleration : ")
		fmt.Scanln(&a)
	
	var v0 float64
		fmt.Print("Please enter velocity : ")
		fmt.Scanln(&v0)
	
	var s0 float64
		fmt.Print("Please enter initial displacement : ")
		fmt.Scanln(&s0)
	
	var t float64
		fmt.Print("Please enter time (t) : ")
		fmt.Scanln(&t)
	
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64{
	return func(t float64) float64 {
		return 1/2*a*t*t+v0*t+s0
	}
}