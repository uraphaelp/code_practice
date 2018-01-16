package movezeroes

import (
	"fmt"
	"testing"
)

type testBed []int

var testcase = []testBed{
	testBed{0, 1, 0, 3, 12},
	testBed{0, 0, 0, 0, 1},
	testBed{0, 0, 0},
	testBed{1, 2, 3, 4},
	testBed{},
}

func TestMoveZeroes(t *testing.T) {
	for _, v := range testcase {
		result := MoveZeroes(v)
		fmt.Println(result)
	}
}
