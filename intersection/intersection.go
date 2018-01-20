package intersection

/*import (
	""
)*/

func Intersection(nums1 []int, nums2 []int) []int {
	var result []int
	map1:=make(map[int]int)
	map2:=make(map[int]int)
	for i:=0; i<len(nums1); i++ {
		map1[nums1[i]]++
	}
	for i:=0; i<len(nums2); i++ {
		//1st version: if map1[nums2[i]]!=0 && map2[nums2[]]?? map1{
		if map2[nums2[i]]<map1[nums2[i]] {
			result=append(result, nums2[i])
			map2[nums2[i]]++
			//确保之后重复的元素不再添加进数组
			//1st version: map1[nums2[i]]=0
		}
	}
	return  result
}