func isMonotic(word string, isUpper bool) bool {
    for _, letter := range word {
        if isUpper && !('A' <= letter && letter <= 'Z') {
            return false
        }
        if !isUpper && !('a' <= letter && letter <= 'z') {
            return false
        }
    }

    return true
}

func detectCapitalUse(word string) bool {
    if 'a' <= word[0] && word[0] <= 'z' {
        return isMonotic(word[1:], false)
    }
    return isMonotic(word[1:], false) || isMonotic(word[1:], true)
}