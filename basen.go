package basen

import "fmt"

type BaseN struct {
    Charset []byte
    BaseNDec map[string]int
    DecBaseN map[int]string
    Base int
}

func NewBaseN(charset []byte, base int) (*BaseN) {
    basendec := make(map[string]int)
    decbasen := make(map[int]string)

    for pos, char := range charset {
        basendec[string(char)] = pos
        decbasen[pos] = string(char)
    }

    return &BaseN {
        Charset: charset
        Base: base,
        BaseNDec: basendec,
        DecBaseN: decbasen,
    }
}

// add byte arrays of base n
func (basen *BaseN) Add(augend []byte, addend []byte) {
    var long []byte
    var short []byte
    // designate long / short quantity
    // reverse byte arrays
    if len(augend) > len(addend) {
        long = reverse(augend)
        short = reverse(addend)
    } else {
        long = reverse(addend)
        short = reverse(augend)
    }

    fmt.Printf("%s + %s", long, short)

    for dec, char := range basen.DecBaseN {
        fmt.Print(dec)
        fmt.Println(" - " + char)
    }

    // for dec := basen.Base - 1

}

// reverse byte array
func reverse(a []byte) []byte {
    length := len(a)
    reversed := make([]byte, length)
    pos := 0
    for i := length - 1; i >= 0; i-- {
        reversed[pos] = a[i]
        pos++
    }
    return reversed
}

// func (basen *BaseN) Add(augend []byte, addend []byte) []byte {
// func (basen *BaseN) Add() {
    // var long []byte
    // var short []byte
    //
    // if len(augend) > len(addend) {
    //     long = augend
    //     short = addend
    // } else {
    //     long = addend
    //     short = augend
    // }
    // resultant := make([]byte, len(long))
    // remainder := 0
    // carry := 0
    // var shortDec int

    // fmt.Printf("%+v", basen)
    // for dec := basen.Base - 1; dec >= 0; dec-- {
    //     fmt.Println(dec)
    // }



    // //
    // // fmt.Printf("Long: %s\nShort: %s\n", long, short)
    // for pos, char := range long {
    //     longDec := base62dec[char]
    //     if pos >= len(short) {
    //         shortDec = 0
    //     } else {
    //         shortDec = base62dec[short[pos]]
    //     }
    //
    //     singleSum := longDec + shortDec + carry
    //     remainder = 0
    //     carry = 0
    //     if singleSum >= 62 {
    //         carry = 1
    //         remainder = singleSum - 62
    //         resultant[pos] = decbase62[remainder]
    //     } else {
    //         resultant[pos] = decbase62[singleSum]
    //     }
    // }
    // if carry == 1 {
    //     r := make([]byte, len(resultant) + 1)
    //     for pos, char := range resultant {
    //         r[pos] = char
    //     }
    //     r[len(resultant)] = decbase62[1]
    //     resultant = r
    // }
    // resultant = reverse(resultant)
    // return resultant
// }
