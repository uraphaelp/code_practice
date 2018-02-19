package disappearnum

import (
	"testing"
	"fmt"
)

func TestFinddisappear(t *testing.T) {
	testcase:=[][]int{{4,3,2,7,8,2,3,1}, {1,5,4,4,1}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Finddisappear(testcase[i]))
	}
}
