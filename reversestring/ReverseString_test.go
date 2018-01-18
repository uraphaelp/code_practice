package reversestring

import (
	"testing"
	"fmt"
)


func TestReverseString(t *testing.T) {
	var s =[]string{/*"", " ", "a", "abc",*/ "hello, world!"}
	for _, v:=range s {
		fmt.Printf("%s\n", ReverseString(v))
	}
}
