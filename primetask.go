package main

import (
	"fmt"
	//"math"
)

func prime() {
	count := 0
	var primecount int
	fmt.Println("prime numbers between 1 to 1000 are")
	for i := 2; i <= 1000; i++ {
		primecount = 0

		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				primecount = 1
				break
			}
		}
		if primecount == 0 && i != 1 {

			fmt.Println("the prime numbers are=", i)
			count++
			//fmt.Println(count)

		}
		//fmt.Println(count)

	}
	fmt.Println("the total primecount between 1 to 1000=", count)

}

func main() {
	// c := make(chan int)
	prime()

}
