func pickGifts(gifts []int, k int) int64 {
    for j := 1; j <= k; j++ {
        maxIdx := 0
        for i := 1; i < len(gifts); i++ {
            if gifts[i] > gifts[maxIdx] {
                maxIdx = i
            }
        }

        gifts[maxIdx] = int(math.Floor(math.Sqrt(float64(gifts[maxIdx]))))
    }
    var sum int64
    for _, d := range gifts {
        sum += int64(d)
    }

    return sum
}
