# find the minimum flip required to get the desired output
# remember the way to convert a decimal to binary number
# using,
# r = r % 10
# n = n / 10
# thes reminder value will be the binary representation of given decimal number..do the operation on it...

func minFlips(a int, b int, c int) int {
    acutalRes := a | b
    flip := 0
    if acutalRes != c {
        for a > 0 || b > 0 || c > 0{
             ra, rb, rc := a%2, b%2, c%2
             a, b, c = a/2, b/2, c/2
            if ! (ra | rb == rc) {
                flip++
                if ra == 1 && rb == 1 && rc == 0{
                    flip++
                }
            }
            
        }
    }
    return flip
    
}
