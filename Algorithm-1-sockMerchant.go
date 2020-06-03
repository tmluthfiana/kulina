package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var n, sock, result int
	fmt.Fscanf(stdin, "%d", &n)
	socks := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &sock)
		socks[sock]++
	}
	for _, c := range socks {
		result += (c / 2)
	}
	fmt.Fprintln(stdout, result)
}

/* another method without input*/

/*
func sockMerchant(socks []int) (result int) {
	result = 0
	mapSocks := make(map[int]int)
	for _, sock := range socks {
		mapSocks[sock]++
	}

	for _, val := range mapSocks {
	   result += (val / 2)
	}
	return result
}

func main() {
	socks := []int{10,20,20,10,10,30,50,10,20}
	fmt.Println(sockMerchant(socks))
}
*/
