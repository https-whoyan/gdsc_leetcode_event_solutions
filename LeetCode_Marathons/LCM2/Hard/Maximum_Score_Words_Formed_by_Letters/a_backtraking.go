var globalWords []string
var globalScoreOfWords []int

func copyMap(originMap map[byte]int) map[byte]int {
    mapToReturn := make(map[byte]int)

    for key, value := range originMap {
        mapToReturn[key] = value
    }

    return mapToReturn
}

func countAns(indexWord int, lettersMap map[byte]int) int {
    if indexWord == len(globalWords) {
        return 0
    }
    
    ansWithoutThisWord := countAns(indexWord + 1, copyMap(lettersMap))

    for _, letter := range globalWords[indexWord] {
        if lettersMap[byte(letter)] == 0 {
            return ansWithoutThisWord
        }
        lettersMap[byte(letter)] -= 1
    }

    ansWithThisWord := countAns(indexWord + 1, copyMap(lettersMap)) + globalScoreOfWords[indexWord]
    if ansWithThisWord > ansWithoutThisWord {
        return ansWithThisWord
    } else {
        return ansWithoutThisWord
    }
}

func maxScoreWords(words []string, letters []byte, score []int) int {
    globalWords = words
    globalScoreOfWords = make([]int, len(words))

    mapOfLetters := make(map[byte]int)
    for _, letter := range letters {
        mapOfLetters[byte(letter)] += 1
    }

    for i := 0; i < len(words); i++ {
        for _, letter := range(words[i]) {
            globalScoreOfWords[i] += score[byte(letter) - byte('a')]
        }
    }

    return countAns(0, mapOfLetters)
}