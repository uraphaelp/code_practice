func Search(nums []int, target int) int {
	if len(nums)==0 {
		return -1
	}
	var (
		roindex, realindex, left int
		right=len(nums)-1
	)
	roindex=Bisearchrotate(nums)
	for left<=right {
		mid:=(left+right)/2
    //定义realindex，整个数组查找
		realindex=(mid+roindex)%(len(nums))
		if nums[realindex]==target {
			return realindex
		} else if nums[realindex]<target {
			left=mid+1
		} else {
			right=mid-1
		}
	}
	return -1
}


/* 步骤3的方法
func Binarysearch(nums []int, target int) int {
	var (
		left int
		right=len(nums)-1
	)
	for left<=right {
		mid:=(left+right)/2
		if target<nums[mid] {
			right=mid-1
		} else if target>nums[mid] {
			left=mid+1
		} else {
			return mid
		}
	}
	return -1
}
*/

func Bisearchrotate(nums []int) int {
	var (
		lo int
		hi=len(nums)-1
	)
	for lo<hi {
		mid:=(lo+hi)/2
		if nums[mid]>nums[hi] {
			lo=mid+1
		} else {
			hi=mid
		}
	}
	return lo
}
