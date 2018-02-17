package setmismatch

import (
	"testing"
	"fmt"
)

func TestFinderrornums(t *testing.T) {
	testcase:=[][]int{{1,2,2,4}, {4,2,1,2}, {1,1}, {3,6,4,1,5,9,2,10,7,7}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Finderrornums(testcase[i]))
		fmt.Printf("\n")
	}
}
