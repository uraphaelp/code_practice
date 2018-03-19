//O(n^3)
func Foursum(nums []int, target int) [][]int {
	var (
		res [][]int
		sum int
		tmp []int
	)
	if len(nums)<4 {
		return res
	}
	sort.Ints(nums)
  //stick to one element
	for i:=0; i<len(nums)-3; i++ {
		if i>0 && nums[i]==nums[i-1] {
			continue //the same element
		}
		if nums[i]*4>target {
			break //to big
		}
		if nums[i]+3*nums[len(nums)-1]<target {
			continue //to small
		}
		for j:=i+1; j<len(nums)-2; j++ {
			if j>i+1 && nums[j]==nums[j-1] {
				continue
			}
			if nums[j]*3>target-nums[i] {
				break
			}
			if nums[j]+2*nums[len(nums)-1]<target-nums[i] {
				continue
			}
      //use two "pointer"
			left:=j+1
			right:=len(nums)-1
			for left<right {
				sum=nums[i]+nums[j]+nums[left]+nums[right]
				if sum==target {
					tmp=append(tmp, nums[i], nums[j], nums[left], nums[right])
					res=append(res, tmp)
					tmp=[]int{}
					for left<right && nums[left]==nums[left+1] {
						left++
					}
					for left<right && nums[right]==nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum<target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
// 通过增加 if 条件判决的种类，降低运行时间
