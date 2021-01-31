package _5string

import (
	"fmt"
	"testing"
)

func TestBFSearch(t *testing.T){
	main := "abcd227fac"
	pattern := "ac"
	fmt.Println(bfSearch(main, pattern))
}