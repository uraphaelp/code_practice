package thirdmaxnum

import (
	"testing"
	"fmt"
)

func TestThirdMax(t *testing.T) {
	testcase:=[][]int{{1,2}, {3,2,1}, {2,1}, {2,1,3}, {2,2,3,1}, {1,2,-2147483648}, {1,1,2}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(ThirdMax(testcase[i]))
	}
}

//OUTPUT: 2, 1, 2, 1, 1, -2147483648, 2
