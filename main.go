package main

import "fmt"

func main() {
	f := []float64{6, 6.1, 6.5, 6.9}
	for i := range f {
		fmt.Println(int(f[i]))
	}
}
