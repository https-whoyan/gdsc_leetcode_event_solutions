func isSubsequence(s string, t string) bool {
    indexS, indexT := 0, 0

    for indexS != len(s) && indexT != len(t) {
        if (s[indexS] == t[indexT]) {
            indexS++
        }
        indexT++
    }

    return indexS == len(s)
}