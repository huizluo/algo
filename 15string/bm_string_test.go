package _5string

import (
	"fmt"
	"testing"
)

func TestBMSearch(t *testing.T) {
	main := "abcacabcbcabcabc"
	pattern := "cabcab"

	fmt.Println(bmSearch(main, pattern))
}
