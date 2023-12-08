//GCDOfStrings finds the greatest common divisor of strings
//find the GCD based on string length
//get the string value based on it

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

    gcdLength := findGCDLength(len(str1), len(str2))
    fmt.Println("gcdLenght:", gcdLength)

    return str1[:gcdLength]
}

//keep updating a and b till it become 0
func findGCDLength(a,b int)int{
    for b != 0{
        a, b = b, a%b
    }
    return a
}
	
