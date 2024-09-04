package util

import (
	"fmt"
	"strings"
)

/// entry function for encoding process
/// argument: fileContent - file content to compress
func Encode(fileContent []byte) []byte {
	/// character -> frequency
	var freqCount map[byte]int = make(map[byte]int)
	/// bit code  -> character
	var charCodes map[string]string = make(map[string]string)
	for _, i := range fileContent {
		freqCount[i]++
	}
	var rootNode *TreeNode = BuildTree(freqCount, charCodes)
	/// encodings will be stored in charCodes
	rootNode.EncodeNode(charCodes, "")
	var encodedString string = coreEncoding(fileContent, charCodes)
	var encodedFileHeader string
	for k, v := range charCodes {
		encodedFileHeader = encodedFileHeader + fmt.Sprintf("%s%s%s%s", k, KeyValueSeparator, v, KeyValuePairSeparator)
	}
	encodedFileHeader = strings.TrimRight(encodedFileHeader, KeyValuePairSeparator) + HeaderBodySeparator
	var encodedFileContent []byte = []byte(encodedFileHeader)
	encodedFileContent = append(encodedFileContent, StringToBits(encodedString)...)
	return encodedFileContent
}

/// core encoding function
/// argument: fileContent - actual file content to encode
/// argument: charCodes - map<huffman encoded bit sequence, character>
func coreEncoding(fileContent []byte, charCodes map[string]string) string {
	var output string
	/// map<character, huffman encoded bit sequence>
	var reverseCharCodes map[string]string = make(map[string]string)
	for k, v := range charCodes {
		reverseCharCodes[v] = k
	}
	for _, ch := range fileContent {
		output = output + reverseCharCodes[string(ch)]
	}
	return output
}
