//给出一个 m*n 全0矩阵；以及一个操作数组：每行只有2个值，每行代表1次操作；每次操作对矩阵中line<操作数组中的第1个元素，column<操作数组中的第2个元素+1
//如：ops=[[2, 2], [3, 3]]代表：首次操作对头2行2列的元素+1；第二次操作对头3行3列的元素+1

func Maxcount(m int, n int, ops [][]int) int {
	var (
		arr [][]int
		count int
	)
	for i:=0; i<len(ops); i++ {
		ref:=ops[i]
		for j:=0; j<ref[0]; j++ {
			for k:=0; k<ref[1]; k++ {
				arr[j][k]++
			}
		}
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if arr[i][j]==arr[0][0] {
				count++
			}
		}
	}
	return count
}

func Maxcount2(m int, n int, ops [][]int) int {
	for i:=0; i<len(ops); i++ {
		m=Findmin(m, ops[i][0])
		n=Findmin(n, ops[i][1])
	}
	return m*n
}

func Findmin(x,y int) int {
	if x>=y {
		return y
	} else {
		return x
	}
}
