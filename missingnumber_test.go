package missingnumber

import (
	"fmt"
	"testing"
)

type testBed []int

var testcase = []testBed{
	testBed{3, 6, 2, 0, 7, 1, 5},
	testBed{3, 1, 0},
	testBed{5, 1, 3, 7, 8, 9, 2, 4, 6},
}

func TestMissingNumber(t *testing.T) {
	for _, v := range testcase {
		result := MissingNumber(v)
		fmt.Println(result)
	}
}
