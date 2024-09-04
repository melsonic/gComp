package util

import (
	"strings"
)

func Decode(input []byte) string {
	var fileContent string = string(input)
	// fileContent structure -> <Header<\n>Encoding>
	var fileContentSlice []string = strings.Split(fileContent, HeaderBodySeparator)
	var charCodes map[string]string = make(map[string]string)
	var headerKV []string = strings.Split(fileContentSlice[0], KeyValuePairSeparator)
	for _, kv := range headerKV {
		kvSlice := strings.Split(kv, KeyValueSeparator)
		charCodes[kvSlice[0]] = kvSlice[1]
	}
	var contentToDecode string = BitsToString([]byte(fileContentSlice[1]))
	return DecodeUtil(contentToDecode, charCodes)
}

// / encoded bit string, character codes(bit string -> character)
func DecodeUtil(encodedString string, charCodes map[string]string) string {
	var output string
	var curr string
	for _, ch := range encodedString {
		if charCodes[curr] != "" {
			output = output + charCodes[curr]
			curr = string(ch)
		} else {
			curr = curr + string(ch)
		}
	}
	output = output + charCodes[curr]
	return output
}
