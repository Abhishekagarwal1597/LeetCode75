func isSubsequence(s string, t string) bool {
   
    if len(t) == 0  && len(s) != 0 {
        return false
    } else if len(s) == 0 && len(t) != 0 {
        return true
    } else if len(s) == 0 && len(t) == 0 {
        return true
    } else {
    
        j := 0
        for i:=0; i < len(t); i++{
            if len(s) > j && t[i] == s[j] {
                j++
            }
        }

        if j == len(s) {
            return true
        } else{
            return false
        }
    }
    
    OR Use the below code ...
    Where time complexity is more optimized..
  
    s_i, t_i := 0, 0
    
    for s_i < len(s) && t_i < len(t) {
        if s[s_i] == t[t_i] { s_i++ }
        t_i++
    }
    return s_i == len(s)
         
}
