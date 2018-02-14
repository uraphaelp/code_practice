func TestCheckperfect(t *testing.T) {
	testcase:=[]int{0, 1, 2, 4, 5, 6, 8, 11, 97, 100}
	for i:=0; i<len(testcase); i++ {
		fmt.Println(Checkperfect(testcase[i]))
	}
}
