package main

import "fmt"

func main() {
	fmt.Println("Hello World in Go");

	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.1426
	x,y := 14,15

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(pi);
	fmt.Println(x);
	fmt.Println(y, x);
	fmt.Printf("%.3f \n", pi)

	for i := 1; i<=10; i++ {
		fmt.Println(i);
	} 
	
	n := 1;

	for n<=10 {
		fmt.Println(n);
		n++
	}
}
