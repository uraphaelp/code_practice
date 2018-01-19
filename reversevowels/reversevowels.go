package reversevowels

//import ""

func Reversevowels(s string) string {
	var count []int
	var sbyte  =[]byte(s)
	for i:=0; i<len(s); i++ {
		switch sbyte[i] {
		case 65:
			count=append(count, i)
		case 97:
			count=append(count, i)
		case 69:
			count=append(count, i)
		case 101:
			count=append(count, i)
		case 73:
			count=append(count, i)
		case 105:
			count=append(count, i)
		case 79:
			count=append(count, i)
		case 111:
			count=append(count, i)
		case 85:
			count=append(count, i)
		case 117:
			count=append(count, i)
		}
	}
	for i:=0; i<len(count)/2; i++ {
		tmp:=sbyte[count[i]]
		sbyte[count[i]]=sbyte[count[len(count)-i-1]]
		sbyte[count[len(count)-i-1]]=tmp
	}
	return string(sbyte)
}
