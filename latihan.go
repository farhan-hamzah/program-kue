package main

import "fmt"

func mesinkue(terigu int, gula int, mentega int, telur int) int {
    var kue int
    kue = 0
    for terigu >= 20 && gula >= 5 && mentega >= 6 && telur >= 1 {
        terigu -= 20
        gula -= 5
        mentega -= 6
        telur--
        kue++
    }
    return kue
}

func main() {
    var t, g, m, te int
    fmt.Scan(&t, &g, &m, &te)
    jumlahKue := mesinkue(t, g, m, te)
    fmt.Println(jumlahKue)
}