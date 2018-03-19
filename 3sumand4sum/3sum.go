//找出给定数组 nums 中，对应给定 target， 任意三个元素值之和为 target 的组合
import (
	"sort"
)

type Arrinf []int

func (arr Arrinf) Len() int {
	return len(arr)
}
func (arr Arrinf) Less(i, j int) bool {
	return arr[i]<arr[j]
}
func (arr Arrinf) Swap(i, j int) {
	arr[i], arr[j]=arr[j], arr[i]
}
//O(n^2)
func Threesum(nums Arrinf, target int) [][]int {
	//need to be sorted
	if !sort.IsSorted(nums) { //apply Issorted must call interface, but just sort.Ints would not need that
		sort.Sort(nums)
	}
	var res [][]int
	for i:=0; i<len(nums)-2; i++ {
		//every time pick a "target", stick to one element
		if i>0 && nums[i]==nums[i-1] {
			continue
		}
		//use two "pointers"
		left:=i+1
		right:=len(nums)-1
		tmptar:=target-nums[i]
		var tmp Arrinf
		for left<right {
			tmp=[]int{}
			if tmptar==nums[left]+nums[right] {
				tmp=append(tmp, nums[i], nums[left], nums[right])
				left++
				right--
				for left<right && nums[left]==nums[left-1] {
					left++
				}
				for left<right && nums[right]==nums[right+1] {
					right--
				}
			} else if tmptar<nums[left]+nums[right] {
				right--
			} else {
				left++
			}
			if len(tmp)!=0 {
				res=append(res, tmp)
			}
		}
	}
	return res
}
