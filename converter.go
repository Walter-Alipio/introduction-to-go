package main

import "fmt"

var k float32 = 373.2

func main(){
	c := int32(k - 273)
	fmt.Printf("O ponto de ebolição da água em Kelvin é : %v e em Celcius é: %v", k, c)
}