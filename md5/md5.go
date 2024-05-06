package md5

import (
	"encoding/hex"
)

var (
	s = make([]uint32, 64)
	k = make([]uint32, 64)

	i uint32
)

var (
	A uint32 = 0x67452301
	B uint32 = 0xefcdab89
	C uint32 = 0x98badcfe
	D uint32 = 0x10325476
)

func Init() {
	s = []uint32{7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21}

	// fmt.Printf("S: %+v\n", s)

	// for i := 0; i < 64; i++ {
	// 	// k[i] = uint32(math.Floor(math.Pow(float64(2), float64(32)) * math.Abs(float64(math.Sin(float64(i + 1))))))
	// 	k[i] = uint32(math.Floor(math.Pow(float64(2), float64(32)) * math.Abs(float64(math.Sin(float64(i+1))))))
	// }

	k = []uint32{0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
		0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
		0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
		0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
		0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
		0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
		0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
		0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
		0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391}
	// fmt.Printf("K: %+v\n", k)
}

func Md5(message string) string {

	paddedMessage := Pad(message)
	// fmt.Println(paddedMessage)

	for i := 0; i < len(paddedMessage); i += 64 {

		block := paddedMessage[i : i+64]

		a, b, c, d := A, B, C, D
		words := BreakIntoWords(block)
		var j uint32
		for j = 0; j < 64; j++ {
			var f, g uint32

			if j < 16 {
				f = (b & c) | ((^b) & d)
				g = j
			} else if j < 32 {
				f = (d & b) | ((^d) & c)
				g = (5*j + 1) % 16
			} else if j < 48 {
				f = b ^ c ^ d
				g = uint32(3*j+5) % 16
			} else {
				f = c ^ (b | (^d))
				g = uint32(7*j) % 16
			}

			f = f + a + k[j] + words[g]
			a = d
			d = c
			c = b
			b = b + leftRotate(f, s[j])
		}

		A += a
		B += b
		C += c
		D += d

	}

	hash := []byte{byte(A), byte(A >> 8), byte(A >> 16), byte(A >> 24),
		byte(B), byte(B >> 8), byte(B >> 16), byte(B >> 24),
		byte(C), byte(C >> 8), byte(C >> 16), byte(C >> 24),
		byte(D), byte(D >> 8), byte(D >> 16), byte(D >> 24)}

	return hex.EncodeToString(hash)
}

func Pad(message string) []byte {

	msg := []byte(message)
	length := len(msg)
	bitLength := uint64(length) * 8 // length

	// append with '1' bit
	msg = append(msg, 0x80)
	// fmt.Printf("Msg: %s\nLength: %d\n, bitLength: %d\n", msg, length, bitLength)
	// // Append with '0's
	for (len(msg)+8)%64 != 0 {
		msg = append(msg, 0x00)
	}

	// Append length in little endian format
	for i := 0; i < 8; i++ {
		msg = append(msg, byte(bitLength>>(8*i)))
	}

	return msg
}

func BreakIntoWords(block []byte) []uint32 {
	words := make([]uint32, 16)

	for i := 0; i < 16; i++ {
		// words[i] = binary.LittleEndian.Uint32(block[i*4 : (i+1)*4])
		word := uint32(block[i*4]) | uint32(block[i*4+1])<<8 | uint32(block[i*4+2])<<16 | uint32(block[i*4+3])<<24
		words[i] = word
	}

	return words
}

func leftRotate(x, c uint32) uint32 {

	return (x << c) | (x >> (32 - c))
}
