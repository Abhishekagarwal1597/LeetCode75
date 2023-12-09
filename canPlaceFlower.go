# to place flower at alternative postion

    func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0

    for i := 0; i < len(flowerbed); i += 1 {
        prev := 0
        next := 0

        if flowerbed[i] == 0 {
            if i > 0 {
            prev = flowerbed[i-1]
        }
         
            if i < len(flowerbed)-1 {
                next = flowerbed[i+1]
            }

            if prev == 0 && next == 0 {
                count++
                i++ // Skip the next position since you've already planted a flower
                
            }
        }
    }

    return count >= n
}

   
