#using brute force apporch
#memory 2.9mb
#time 2ms

func mergeAlternately(word1 string, word2 string) string {

    if word1 == "" {
		return word2
	}

	if word2 == "" {
		return word1
	}

    var word3 string
    for i,j := 0, 0; i < len(word1) && j < len(word2); i,j = i+1, j+1{
        word3 += string(word1[i]) 
        word3 += string(word2[j])
    }

    if len(word1) > len(word2) {
        for i:=len(word2);i<len(word1);i++{
            word3 += string(word1[i])
        }
    }

    if len(word2) > len(word1) {
        for i:=len(word1);i<len(word2); i++{
            word3 += string(word2[i])
        }
    }

    return word3

}
