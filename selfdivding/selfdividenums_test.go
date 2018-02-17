package selfdividenums

import (
	"testing"
	"fmt"
)

func TestSelfdividingnums(t *testing.T) {
	testcase:=[]int{1, 4, 5, 10, 15, 22}
	for i:=0; i<len(testcase); i++ {
		for j:=0; j<len(testcase); j++ {
			fmt.Println(Selfdividingnums(testcase[i], testcase[j]))
		}
	}
}
