// use string manuipulation function, use field instead of Split and trim
// append work with slice
// use 2 pointer approch 

func reverseWords(s string) string {
    
    // stringArray := strings.Split(strings.TrimSpace(s), " ")
    stringArray := strings.Fields(s)
    fmt.Println("stringArray:", stringArray)
    var newStringArray []string
    for i:=len(stringArray)-1 ; i>=0; i--{
        newStringArray = append(newStringArray, stringArray[i])
    }
    fmt.Println("newStringArray:", newStringArray)
    // re := regexp.MustCompile(`\s+`)
    return strings.Join(newStringArray, " ")
    // newUpdatedString := strings.Join(newStringArray, " ")
    // return re.ReplaceAllString(newUpdatedString, " ")

}

or 

func reverseWords(s string) string {
    // Split the input string into words
    words := strings.Fields(s)

    // Initialize pointers for the word reversal
    low, high := 0, len(words)-1

    // Swap words using the two-pointer approach
    for low < high {
        words[low], words[high] = words[high], words[low]
        low++
        high--
    }

    // Join the reversed words back into a single string and trim spaces
    return strings.TrimSpace(strings.Join(words, " "))
}

