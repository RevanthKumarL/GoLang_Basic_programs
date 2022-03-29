
package main
import "fmt"

func fact(n int) int {
if n == 0 {
	return 1
}
	return n * fact(n - 1)
}

func main() {
	fmt.Println(fact(7))
	var fib func(n int) int  // in case of recursions with closures, closures should be explicitly declared using a var


	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n - 1) + fib(n - 2) // fib series: 0 1 1 2 3 5 8 13
	}
	fmt.Println(fib(7))
}