package missingnumber

//import "fmt"

func MissingNumber(nums []int) int {
	var practicalsum, realsum, l int
	l = len(nums)
	for i := 0; i < l; i++ {
		practicalsum += nums[i]
	}
	realsum = (1 + l) * l / 2
	return (realsum - practicalsum)
}

/*
func main() {
	arr1 := []int{1, 3, 2, 0, 6, 4}
	arr2 := []int{1, 3, 4, 2, 6, 7, 5}
	fmt.Println(missingNumber(arr1))
	fmt.Println(missingNumber(arr2))
}
*/
