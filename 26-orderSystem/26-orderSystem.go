func convertToTitle(n int) string {
    var result string
    for n!=0 {
        n--
        result=string('A' + n%26)+result
        n/=26
    }
    return result
}
// see A as 0; Z as 25 instead of A as 1, Z as 26
