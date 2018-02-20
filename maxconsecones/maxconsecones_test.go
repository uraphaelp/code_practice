package maxconsecone

import (
	"testing"
	"fmt"
)

func TestFindmaxconsecutiveones(t *testing.T) {
	testcase:=[][]int{{1,1,0,1,1,1}, {1,0,1,0,0,0,1,0,1,1,1,0,1,0}, {1,0,1,0,0,1}, {},{0,0,0}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Findmaxconsecutiveones(testcase[i]))
	}
}

//OUTPUT: 3, 3, 1, 0, 0
