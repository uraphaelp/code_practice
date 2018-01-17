package powerofthree

import (
	"testing"
	"fmt"
)

func TestIsPowerofThree(t *testing.T) {
	var testcase=[]int{0,1,2,3,6,9,12,15,18,27,243,256}
	for _, v:=range testcase {
		result:=IsPowerofThree(v)
		fmt.Printf("%b\n", result)
	}
}
