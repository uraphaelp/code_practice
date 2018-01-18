package reversestring

//import ""

func ReverseString(s string) string  {
	var tmp byte
	var l int=len(s)
	var sbyte=[]byte(s)
	for i:=0; i<l/2; i++ {
		tmp=sbyte[i]
		sbyte[i]=sbyte[l-i-1]
		sbyte[l-i-1]=tmp
	}
	s=string(sbyte)
	return s
}
