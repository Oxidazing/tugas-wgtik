package main

import "fmt"

func main() {
	var berat, rata float32
	var n int

	fmt.Scanln(&n)

	for i := 0; i != n; i++ {
		fmt.Scanln(&berat)
		berat = berat + berat
	}
	rata = berat / float32(n)
	fmt.Println(rata)
}
