package main
import "fmt"

func vals() (int, int) {  // (int, int) func to return 2 ints
	return 3, 7
}

func main() {
a, b := vals()
fmt.Println(a)
fmt.Println(b)

_, c := vals()  // if at all I'd want a subset of the return value
fmt.Println(c)
//fmt.Println(d)

}