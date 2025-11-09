func addStrings(num1 string, num2 string) string {
    var i, j, carry int
    var res []byte
    N1 := len(num1) - 1
    N2 := len(num2) - 1
    i ,j = N1, N2
    
    for j > -1 || i > -1 || carry > 0 {
        var a, b byte
        
        if i > -1 {
            a = num1[i] - '0'
        }
        if j > -1 {
            b = num2[j] - '0'
        }
        
        i--; j--
        
        res = append([]byte{byte((carry + int(a + b))% 10) + '0'}, res...)
        carry = (carry + int(a + b)) / 10
    }
    
    return string(res)
}