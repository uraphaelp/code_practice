import (
	"testing"
	"fmt"
)

func TestCombinationsum(t *testing.T) {
	testcase:=[][]int{{10,1,2,7,6,1,5}, {6,1,2,3,3,2,5}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Combinationsum(testcase[i], 8))
	}
}
//Output:[[1 1 6] [1 2 5] [1 7] [2 6]]
//[[1 2 2 3] [1 2 5] [2 3 3] [2 6] [3 5]]
