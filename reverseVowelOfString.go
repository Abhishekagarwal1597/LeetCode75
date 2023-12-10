func reverseVowels(s string) string{
    // Convert string to a slice of runes
    chars := []rune(s)
    for i,j:=0,len(chars)-1; i<j; {
        if isVowels(chars[i]) && isVowels(chars[j]){
            chars[i], chars[j] = chars[j], chars[i]
            i++
            j--
        } else if isVowels(chars[i]){
            j--
        } else {
            i++
        }
    }
    
    return string(chars)
}

func isVowels(chars rune)bool{
   return chars == 'a' || chars == 'e' || chars == 'i' || chars == 'o' || chars == 'u' || chars == 'A'|| chars == 'E' || chars == 'I' || chars == 'O' || chars == 'U'
    // return true
    // }else{
    //     return false
    // }
}

// func reverseVowels(s string) string {
//     arr, n := getAllVowels(s)
//     if n == 0 {
//         return s
//     }

//     return reverseVowel(s, arr)
// }


// func getAllVowels(s string)([]rune, int){
//     var vowelArray []rune
//     var n int

//     for i, _ := range []rune(s) {
//        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A'|| s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
//          vowelArray = append(vowelArray, rune(s[i]))
//          n++
//      }
//     }

//     return vowelArray, n
// }

// func reverseVowel(s string, arr []rune)string{
//     fmt.Println("arr:", string(arr))
//     fmt.Println("s:", s)

//     var testrune []rune
//    j :=0

//     for i, _ := range []rune(s) {
//         if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A'|| s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U'{
//             testrune = append(testrune, arr[len(arr)-j -1])
//             j++
//      } else{
//         testrune = append(testrune, rune(s[i]))
//      }
//     }

//     return string(testrune)
// }
