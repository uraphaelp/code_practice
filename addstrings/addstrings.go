package addstrings

import "strconv"

func Addstrings(num1 string, num2 string) string {
	var (
		i=len(num1)-1
		j=len(num2)-1
		bit int
		result string
	)
	for i>=0 || j>=0 || bit==1 {
		var x, y int
		if i>=0 {
			x=int(num1[i]-'0')
		} else {
			x=0
		}
		if j>=0 {
			y=int(num2[j]-'0')
		} else {
			y=0
		}
		result+=strconv.Itoa((x+y+bit)%10)
		bit=(x+y+bit)/10
		i--
		j--
	}
	return Reverse(result)
}

func Reverse(s string) string {
	var result string
	for i:=len(s)-1; i>=0; i-- {
		result+=string(s[i])
	}
	return result
}
