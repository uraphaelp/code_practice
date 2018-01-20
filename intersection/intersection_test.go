package intersection

import (
	"testing"
	"fmt"
)

type testbed []int

func TestIntersection(t *testing.T) {
	var testcase1=[]testbed{{1,2,2,1}, {}, {1}, {1,1,3,2,2,1,4,9,0,7,5}, {1,2,3,4,5,6,100,72}}
	var testcase2=[]testbed{{2,2}, {}, {2}, {1,6,0,5}, {0,0,0,1,23,46,72}}
	for i:=0; i<len(testcase1); i++ {
		for j:=0; j<len(testcase2); j++ {
			fmt.Println(Intersection(testcase1[i], testcase2[j]))
		}
	}
}
