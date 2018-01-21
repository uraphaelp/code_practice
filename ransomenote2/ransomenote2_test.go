package ransomenote2

import (
	"testing"
	"fmt"
)

func TestCanconstruct(t *testing.T) {
	testcase1:=[]string{"", "a", "ab", "abaa", "bdfafa"}
	testcase2:=[]string{"a", "aab", "aaabbbff", "avabaabfbafaa"}
	for i:=0; i<len(testcase1); i++ {
		for j:=0; j<len(testcase2); j++ {
			fmt.Printf("%t\n", Canconstruct(testcase1[i], testcase2[j]))
		}
	}
}
