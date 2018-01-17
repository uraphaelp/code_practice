package powerofthree

//import ""

func IsPowerofThree(n int) bool {
	for n>=3 {
		if n%3!=0 {
			return false
		}
		n/=3
	}
	if n!=1 {
		return false
	}
	return true
}
