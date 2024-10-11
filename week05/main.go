package main
import "fmt" //C언어의 #include <stdio.h> 역할

func main() {
	//var i int = 55;
	var f float32 = 3.99

	//var i int
	//i = 55
	
	i := 55
	fmt.println(f);
	fmt.printf("i는 %d\n",i);
	fmt.print("i는",i "\n");
	fmt.println("i는",i);
}