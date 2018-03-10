func maxSumOfThreeSubarrays(nums []int, K int) []int {
   var W []int
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i >= K {
			sum -= nums[i-K]
		}
		if i >= K-1 {
			//W[i-K+1] = sum
			W=append(W, sum)
		}
	}
	//find index from left
	//left:=[len(W)]int{}
	var left []int
	best := 0
	for i := 0; i < len(W); i++ {
		if W[i] > W[best] {
			best = i
		}
		//left[i] = best
		left=append(left, best)
	}
	//find index from right
	//right:=[len(W)]int{}
	var right []int
	best= len(W) - 1
	for i := len(W) - 1; i >= 0; i-- {
		if W[i] >= W[best] {
			best = i
		}
		//right[i] = best
		right=append(right, best)
	}
	right=Reversearr(right)
	//find the smallest index and avoid overlapping
	ans := []int{-1, -1, -1}
	for j := K; j < len(W) - K; j++ {
		i := left[j - K]
		k := right[j + K]
		if ans[0] == -1 || W[i] + W[j] + W[k] > W[ans[0]] + W[ans[1]] + W[ans[2]] {
			ans[0] = i
			ans[1] = j
			ans[2] = k
		}
	}
	return ans
}

func Reversearr(nums []int) []int {
	for i:=0; i<len(nums)/2; i++ {
		tmp:=nums[i]
		nums[i]=nums[len(nums)-i-1]
		nums[len(nums)-i-1]=tmp
	}
	return nums
}
