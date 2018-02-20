//给定一个二维数组表示的矩阵以及，行数r、列数c；将原矩阵转换成r*c的新矩阵

func Matrixreshape(nums [][]int, r,c int) [][]int {
	if len(nums)==0 || r*c!=accountofmatrix(nums) {
		return nums
	}
	var (
		res [][]int
		tmp []int
		index=0
	)
	for i:=0; i<len(nums); i++ {
		for j:=0; j<len(nums[0]); j++ {
			tmp=append(tmp, nums[i][j])
		}
	}
	for i:=0; i<r; i++ {
    //注意二维数组的添加元素方法
		row:=[]int{}
		for j:=0; j<c; j++ {
			row=append(row, tmp[index])
			index++
		}
		res=append(res, row)
	}
	return res
}

//求矩阵元素数目
func accountofmatrix(nums [][]int) int {
	var (
		r, c int
		tmp []int
	)
	r=len(nums)
	tmp=nums[0]
	c=len(tmp)
	return r*c
}
