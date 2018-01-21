package ransomnote

import (
	"testing"
	"fmt"
)

func TestCanconstruct(t *testing.T) {
	var testcase1  =[]string{"a", "aa", "ab", "abc"}
	var testcase2  =[]string{"a", "aab", "ababcdgfabdabc", "ba"}
	for i:=0; i<len(testcase1); i++ {
		for j:=0; j<len(testcase2); j++ {
			fmt.Printf("%t\n", Canconstruct(testcase1[i], testcase2[j]))
		}
	}
}
