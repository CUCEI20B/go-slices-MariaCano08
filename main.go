package main

import "fmt"

func main() {
	var si []uint64
	var c, i, d uint64
	fmt.Scan(&c)
	d = c
	for i = 0; i < d; {
		fmt.Scan(&c)
		si = append(si, c)
		i = i + 1
	}
	//fmt.Println(si)
	c = 0
	for i = 0; i < d; {
		c = c + si[i]
		//fmt.Println(c)
		i = i + 1
	}
	fmt.Println(c)

}
