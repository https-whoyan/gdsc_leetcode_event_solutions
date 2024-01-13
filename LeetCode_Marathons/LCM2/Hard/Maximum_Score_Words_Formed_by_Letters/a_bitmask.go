func getMask(number int, maskSize int) string {
    mask := ""
    for number != 0 {
        mask = strconv.Itoa(number % 2) + mask
        number /= 2
    }
    mask = strings.Repeat("0", maskSize - len(mask)) + mask
    return mask
}

func maxScoreWords(words []string, letters []byte, score []int) int {
    n := len(words)
    maskCount := int(math.Pow(float64(2), float64(n)))

    ans, scoreCounter := 0, 0
    mpLetters := make(map[byte]int)

    for _, myRune := range letters  {
        mpLetters[myRune]++
    }

    for i := 1; i < maskCount; i++ {
        mask := getMask(i, n)
        
        mpLettersOfMask := make(map[byte]int)
        for indexOfMask := 0; indexOfMask < n; indexOfMask++ {
            if mask[indexOfMask] == '1' {
                for _, myChr := range words[indexOfMask] {
                    mpLettersOfMask[byte(myChr)]++
                }
            }
        }

        scoreCounter = 0
        for key, value := range mpLettersOfMask {
            value2 := mpLetters[key]
            if value > value2 {
                scoreCounter = 0
                break
            } else {
                scoreCounter += score[ rune(key) - rune('a') ] * value
            }
        }
        
        if scoreCounter > ans {
            ans = scoreCounter
        }
    }

    return ans
}