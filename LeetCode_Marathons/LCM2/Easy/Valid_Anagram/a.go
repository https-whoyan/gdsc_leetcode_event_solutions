func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sBytes := []byte(s)
    tBytes := []byte(t)

    sort.Slice(sBytes, func(i, j int) bool {
        return sBytes[i] < sBytes[j]
    })
    sort.Slice(tBytes, func(i, j int) bool {
        return tBytes[i] < tBytes[j]
    })

    for i := 0; i < len(sBytes); i++ {
        if sBytes[i] != tBytes[i] {
            return false
        }
    }
    return true
}