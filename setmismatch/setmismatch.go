//数组 nums 存储 1-n的数字；其中有一个数重复出现一次，一个数缺失，找出这两个数返回

import (
	"sort"
	"fmt"
)

func Finderrornums(nums []int) []int {
	var (
		result []int
		target int
		l=len(nums)
	)
	sort.Ints(nums)
  //检验排序结果
	fmt.Println(nums)
	count:=nums[0]
	for i:=1; i<l; i++ {
		if nums[i]==nums[i-1] {
			result=append(result, nums[i])
		} else {
      //实际减去重复数的和
			count+=nums[i]
		}
	}
  //正常数组的和
	target=(1+l)*l/2
	result=append(result, target-count)
	return result
}
