package main

import "fmt"

const SIZE = 1999

func hashes(a *[]int) []int {
    xs := []int{0, 0, 0}
    for _, v := range *a {
        xs[0] = xs[0]*137 + v
        xs[1] = xs[1]*69 + v
        xs[2] = xs[2]*545 + v
    }

    hashes := make([]int, 3)
    for i, _ := range xs {
        hashes[i] = xs[i] % SIZE
    }

    return hashes
}

type BloomFilter struct {
    BitArray *[]int
}

func (bf *BloomFilter) add(a *[]int) {
    for _, v := range hashes(a) {
        (*bf.BitArray)[v] = 1
    }
}

func (bf *BloomFilter) query(a *[]int) bool {
    for _, v := range hashes(a) {
        if (*bf.BitArray)[v] == 0 {
            return false
        }
    }
    return true
}

func main() {
    ba := make([]int, 0)
    for i := 0; i < SIZE; i++ {
        ba = append(ba, 0)
    }
    bf := &BloomFilter{&ba}

    ary := []int{1, 2, 3}

    bf.add(&ary)

    if bf.query(&ary) {
        fmt.Println("ok")
    }

    ary2 := []int{1, 2, 3, 4}

    if !bf.query(&ary2) {
        fmt.Println("ng")
    }
}