package basen

type BaseN struct {
    Charset []byte
    BaseNDec map[byte]int
    DecBaseN map[int]byte
    Base int
}

func NewBaseN(charset []byte, base int) (*BaseN) {
    basendec := make(map[byte]int)
    decbasen := make(map[int]byte)

    for pos, char := range charset {
        basendec[char] = pos
        decbasen[pos] = char
    }

    return &BaseN {
        Charset: charset,
        Base: base,
        BaseNDec: basendec,
        DecBaseN: decbasen,
    }
}

// add byte arrays of base n
func (basen *BaseN) Add(augend []byte, addend []byte) []byte {
    // first vars
    var long []byte
    var short []byte
    // designate long / short quantities
    // reverse byte arrays
    if len(augend) > len(addend) {
        long = reverse(augend)
        short = reverse(addend)
    } else {
        long = reverse(addend)
        short = reverse(augend)
    }

    // second vars
    resultant := make([]byte, len(long))
    remainder := 0
    carry := 0

    var sdec int

    // iterate through long by pos and character
    for lpos, lchar := range long {
        ldec := basen.BaseNDec[lchar]
        if lpos >= len(short) {
            sdec = 0
        } else {
            sdec = basen.BaseNDec[short[lpos]]
        }

        singleSum := ldec + sdec + carry
        remainder = 0
        carry = 0
        if singleSum >= basen.Base {
            carry = 1
            remainder = singleSum - basen.Base
            resultant[lpos] = basen.DecBaseN[remainder]
        } else {
            resultant[lpos] = basen.DecBaseN[singleSum]
        }
    }

    if carry == 1 {
        r := make([]byte, len(resultant) + 1)
        for pos, char := range resultant {
            r[pos] = char
        }
        r[len(resultant)] = basen.DecBaseN[1]
        resultant = r
    }
    resultant = reverse(resultant)
    return resultant
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
