package basen

import (
    "testing"
    "fmt"
)

////////////////////////////////////////////////////////////////////////////////
//  This is a test module for BaseN package
//      *   demonstrates the abilities to add in any base given a charset
////////////////////////////////////////////////////////////////////////////////
func TestBaseN (t *testing.T) {
    // new base n struct pointer
    // currently base 16
    bn := NewBaseN([]byte("0123456789abcdef"), 16)


    // what we are looking for
    target := []byte("a0")
    // where we are starting
    current := []byte("00")
    // what we will increment by
    increment := []byte("01")
    // where we will end // if target not found in this block
    end := []byte("f0")

    // iterate / increment until we have either found the target
    // or we have reached the end of the block
    flag := false
    for flag != true {
        // did we find the target
        if string(current) == string(target) {
            fmt.Println("found it")
            fmt.Printf("%s\n", current)
            flag = true
        // have we reached the end
        } else if string(current) == string(end) {
            fmt.Println("reached the end")
            flag = true
        // if not then, increment
        } else {
            current = bn.Add(current, increment)
            fmt.Printf("%s\n",current)
        }
    }
}
