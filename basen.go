package basen

import "fmt"

type BaseN struct {
    Charset []byte
    Base int
}

func NewBaseN(charset []byte, base int) (*BaseN) {
    return &BaseN{
        Charset: charset,
        Base: base,
    }
}

// func reverse(array []byte) []byte {
//     length := len(array)
//     reversed := make([]byte, length)
//     pos := 0
//     for i := length - 1; i >= 0; i-- {
//         reversed[pos] = array[i]
//         pos++
//     }
//     return reversed
// }

// func (basen *BaseN) Add(augend []byte, addend []byte) []byte {
func (basen *BaseN) Add() {
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

    fmt.Printf("%+v", basen)
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
}
