package reversevowels

import (
	"fmt"
	"testing"
)

func TestReversevowels(t *testing.T) {
	var testbed =[]string{"hello", "leetcode", "", " ", "aaa", "bgm", "reverse", "aAA"}
	for i:=0; i<len(testbed); i++ {
		fmt.Printf("%s", Reversevowels(testbed[i]))
	}
}
