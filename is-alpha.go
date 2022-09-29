package exam

func IsAlpha(str rune) bool {
    if (str > 64 && str < 91) || (str > 96 && str < 123) {
        return true
    } else {
        return false
    }
}
