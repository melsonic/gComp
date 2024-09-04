package util

import "fmt"

func PrintOutputWithCaption(caption string, output string) {
	fmt.Printf("%-20s => %s\n\n", caption, output)
}

/// convert the string representation of encoded bits to actual bits
/// actual compression happends here
func StringToBits(bitString string) []byte {
	var output []byte
	for i := 0; i < len(bitString); i += 8 {
		var curr byte = 0
		for j := i; j < min(len(bitString), i+8); j++ {
			if bitString[j] == '1' {
				curr = (curr << 1) | 1
			} else {
				curr = (curr << 1)
			}
		}
		output = append(output, curr)
	}
	return output
}

/// convert []byte(bit sequence) to string representation of bits
/// argument: input - slice of bytes or bit sequence
func BitsToString(input []byte) string {
	var bitString string
	for idx, inputByte := range input {
		var curr string
		if idx == len(input)-1 {
			curr = fmt.Sprintf("%b", inputByte)
		} else {
			curr = fmt.Sprintf("%08b", inputByte)
		}
		bitString = bitString + curr
	}
	return bitString
}
