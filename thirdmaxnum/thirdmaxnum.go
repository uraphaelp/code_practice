//找出数组中第3大的元素，使用o(n)复杂度
//若没有第3大的，则返回最大的；[1,2,2,3]第3大的是1

func ThirdMax(nums []int) int {
	var (
		l=len(nums)
    //±2^31=2147483648是有符号整型的最大值
		max, mid, min = -2147483649, -2147483649, -2147483649
	)
	if l==1 {
		return nums[0]
	}
	if l==2 {
		return Max(nums[0], nums[1])
	}
  //循环里依次扫描，每次扫描后把当前扫描到的元素之前的所有元素：最大存到max, 第2存到mid, 第3存到min
	for i:=0; i<l; i++ {
		if nums[i]>max {
			min=mid
			mid=max
			max=nums[i]
		} else if nums[i]!=max && nums[i]>mid {
			min=mid
			mid=nums[i]
		} else if nums[i]!=max && nums[i]!=mid && nums[i]>min {
			min=nums[i]
		}
	}
  //是否有重复元素，致使并没有第3大的
	if min!=-2147483649 {
		return min
	} else {
		return max
	}
}

func Max(a,b int) int {
	if a>=b {
		return a
	} else {
		return b
	}
}
