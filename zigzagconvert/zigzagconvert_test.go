import (
	"testing"
	"fmt"
)

func TestConvert(t *testing.T) {
	testcase:=[]string{"", "abcdefgh", "aba", "abcdefghijklmnopqrst"}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Convert(testcase[i], 2))
		fmt.Println(Convert(testcase[i], 5))
	}
}
