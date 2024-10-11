package main

import "fmt"

func main() {
	var i int = 13
	var f  float32 := 12.9
	fmt.printf("value i :%d , value f : %f\n" , i, f)
	//fmt.printf("%d * %f = %f\n", i, f, i*f)
	fmt.printf("%d * %f = %f\n", i, f, float32(i)*f)
	fmt.println(reflect.TypeOf(i))
}