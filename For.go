package main
import "fmt"
func main() {
	i :=1
	for  i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <=9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 7; n++ {
		if n%2 == 1 {    // n is divisible by 2
			continue
		}              
		fmt.Println(n)
	}
}