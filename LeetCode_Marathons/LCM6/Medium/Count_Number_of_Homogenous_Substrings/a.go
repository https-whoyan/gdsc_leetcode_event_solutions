func countHomogenous(s string) int {
    var module int64 = 1000000007
    var ans int64 = 0
    var keyRune rune = rune('*')
    s += "*"
    var keyCounter int64 = 0
    for _, myCh := range s {
        if myCh == keyRune { 
            keyCounter++ 
        } else {
            var ansForRune int64
            ansForRune = (keyCounter + 1) * keyCounter / 2;
            ans = (ans + ansForRune) % module;
            keyCounter = 1
            keyRune = myCh
        }
    }

    return int(ans);
}