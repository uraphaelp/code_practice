func Convert(s string, numrows int) string {
	if numrows==1 || len(s)<=numrows {
		return s
	}
	var (
		//pay attention to this definition, or(res [][]byte, overflow occurs
		res=make([][]byte, len(s), len(s))
		row int
		inc=-1
		ret string
	)
	for i:=0; i<len(s); i++ {
		// to the turning point
		if i%(numrows-1)==0 {
			inc*=-1
		}
		tmp:=res[row]
		tmp=append(tmp, s[i])
		res[row]=tmp
		row+=inc
	}
	for i:=0; i<len(res); i++ {
		for j:=0; j<len(res[i]); j++ {
			ret+=string(res[i][j])
		}
	}
	return ret
}
/*
below is how the zigzag works
 */
/*n=numRows
Δ=2n-2    1                           2n-1                         4n-3
Δ=        2                     2n-2  2n                    4n-4   4n-2
Δ=        3               2n-3        2n+1              4n-5       .
Δ=        .           .               .               .            .
Δ=        .       n+2                 .           3n               .
Δ=        n-1 n+1                     3n-3    3n-1                 5n-5
Δ=2n-2    n                           3n-2                         5n-4
*/
