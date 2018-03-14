//定位出字符串中最长回文序列

func Longestpalindrome(s string) string {
	var (
		left, right int
	)
	for i := 0; i < len(s); i++ {
    //use s[i] as center, expand from both sides
		len1 := Expandfromcenter(s, i, i)
    //do not miss this situation: s[i]&s[i+1] maybe identical
		len2 := Expandfromcenter(s, i, i + 1)
		len := Max(len1, len2)
		if len > right - left {
			left = i - (len - 1) / 2
			right= i + len / 2
		}
	}
	return s[left: right + 1]
}

func Expandfromcenter(s string, left, right int) int {
	l:=left
	r:=right
	for l>=0 && r<len(s) && s[l]==s[r] {
		l--
		r++
	}
	return r-l-1
}

func Max(a, b int) int {
	if a>=b {
		return a
	} else {
		return b
	}
}
