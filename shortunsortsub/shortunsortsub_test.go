package shortestunsortsub

import (
	"testing"
	"fmt"
)

func TestFindunsortsub(t *testing.T) {
	testcase:=[][]int{{1,2,3,4}, {1,1,2,3,4,4,5}, {2,6,4,8,10,9,15}, {1,2,3,5,4,6,7,10,9}, {1,5,2,3,4}, {1,2,3,0,10,6,7,8,9,10}, {1,3,2}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Findunsortsub(testcase[i]))
	}
}

//OUTPUT: 0, 0, 5, 6, 4, 9, 2
