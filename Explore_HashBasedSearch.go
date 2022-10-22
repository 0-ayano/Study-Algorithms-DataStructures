package main

import "fmt"

var table [][]int

func loadTable(a *[]int) {
    table = make([][]int, len(*a))
    for _, e := range *a {
        h := hash(e, len(*a))
        if table[h] == nil {
            table[h] = make([]int, 0)
        }
        table[h] = append(table[h], e)
    }
}

func hash(key int, size int) int {
    var h int
    for i := 0; i < key+1; i++ {
        h = (h*137 + i) % size
    }
    return h
}

func search(a *[]int, t int) bool {
    h := hash(t, len(*a))
    list := table[h]
    if list == nil {
        return false
    }
    for _, v := range list {
        if v == t {
            return true
        }
    }
    return false
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    loadTable(&a)
    fmt.Printf("0:%t\n", search(&a, 0))
    fmt.Printf("2:%t\n", search(&a, 2))
    fmt.Printf("12:%t\n", search(&a, 12))
}