func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    firstPrefix := strs[0]

    for j := 1; j < len(strs); j++ {
        if len(firstPrefix) == 0 || len(strs[j]) == 0 {
            return ""
        }
        minLen := min(len(firstPrefix), len(strs[j]))
        var i int
        for i = 0; i < minLen; i++ {
            if firstPrefix[i] != strs[j][i] {
                break
            }
        }
        firstPrefix = firstPrefix[:i]
    }

    return firstPrefix
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
