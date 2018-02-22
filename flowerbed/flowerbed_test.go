package flowerbed

import (
	"testing"
	"fmt"
)

func TestCanplaceflowers(t *testing.T) {
	testcase:=[][]int{{0,0,0,0}, {0,0,1,0,1,0,1,0,0}, {1,0,0,0,1}, {1,0,0,0,0,0,0,0,0,1,0,0,0}, {0,0,1,0,1,0,0,1,0,0,1,0,1,0,0,0,1,0}}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Canplaceflowers(testcase[i], 2))
	}
}

//OUTPUT: t, t, f, t, t
