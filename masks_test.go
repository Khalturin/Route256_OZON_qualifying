package main

import (
	"fmt"
	"testing"
)

func TestMasks(t *testing.T) {
	fmt.Println("385, right: 1 0 1")
	first(385)
	fmt.Println("23, right: 0 1 0")
	first(23)
	fmt.Println("27, right: 3 1 0")
	first(27)
	fmt.Println("30, right: 6 1 0")
	first(30) //16,5 // 14,3
	fmt.Println("47, right: 0 2 0")
	first(47)
	fmt.Println("12, right: 12 0 0")
	first(12)
	fmt.Println("20, right: 0 1 0")
	first(20)
	fmt.Println("21, right: 0 1 0")
	first(21)
}
