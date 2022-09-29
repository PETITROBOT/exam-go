package exam

func PrintElements(tab []int, str []string) string {
    res := ""
    for _, c := range tab {
        if c < len(str) {
            res = res + " "
            res = res + str[c]
        }
        if c > len(str)-1 {
            res = res + " " + "!"
        }
    }
    return res
}
