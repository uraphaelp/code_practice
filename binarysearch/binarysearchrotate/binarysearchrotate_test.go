import (
	"testing"
	"fmt"
)

func TestSearch(t *testing.T) {
	testcase:=[][]int{{7,8,9,0,1,2,3,4,5,6}, {5,6,7,7,8,0,1,2,3,4,5}, {4,5,6,7,0,1,2}, {1}, {2,1}, {}, {1,3}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Search(testcase[i], 1))
		fmt.Println(Search(testcase[i], 3))
	}
}
