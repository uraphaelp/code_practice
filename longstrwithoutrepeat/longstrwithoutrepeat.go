// 定位字符串中不带重复元素的最长子串
//use slide windows(which requires two index, here is i and j)
func Lengthoflongestsub(s string) int {
	var (
		res, max int
		m=make(map[byte]int)
		l=len(s)
		i, j=0, 0
	)
	for i<l && j<l {
		if m[s[j]]==0 {
			m[s[j]]++
			max++
			j++
			if max>res {
				res=max
			}
		} else {
			m[s[i]]--
			i++
			max--
		}
	}
	return res
}
//注意slide window的使用
