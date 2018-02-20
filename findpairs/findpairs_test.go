package findpairs

import (
	"testing"
	"fmt"
)

func TestFindpairs(t *testing.T) {
	testcase1:=[][]int{{3,1,4,1,5}, {1,2,3,4,5}, {1,3,1,5,4}}
	testcase2:=[]int{2,1,0}
	for i:=0; i<len(testcase1); i++ {
		fmt.Println(Findpairs(testcase1[i], testcase2[i]))
	}
}
