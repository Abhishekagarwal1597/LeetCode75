func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, val := range candies{
        if max < val{
            max = val
        }
    }

    var resultArray []bool
    for _, val := range candies{
        if val >= (max - extraCandies){
            resultArray = append(resultArray, true)
        } else {
            resultArray = append(resultArray, false)
        }
    }
    return resultArray
}
