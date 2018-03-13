import (
	"testing"
	"fmt"
)

func TestLengthoflongestsub(t *testing.T) {
	testcase:=[]string{"abcabcbb", "bbbbb", "pwwkew", "", "a", "dvdf"}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Lengthoflongestsub(testcase[i]))
	}
}
//OUTPUT:3, 1, 3, 0, 1, 3
