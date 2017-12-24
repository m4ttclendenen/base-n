package basen

import (
    "testing"
    "fmt"
)

func TestBaseN (t *testing.T) {
    bn := NewBaseN([]byte("0123456789abcdef"), 16)

    current := []byte("00")
    increment := []byte("01")
    end := []byte("f0")
    target := []byte("a0")
    flag := false
    for flag != true {
        if string(current) == string(target) {
            fmt.Println("found it")
            fmt.Printf("%s\n", current)
            flag = true
        } else if string(current) == string(end) {
            fmt.Println("reached the end")
            flag = true
        } else {
            current = bn.Add(current, increment)
            fmt.Printf("%s\n",current)
        }
    }
}
