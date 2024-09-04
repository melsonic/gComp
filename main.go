package main

import (
	"fmt"
	"log"
	"os"

	"github.com/melsonic/gcomp/util"
)

func main() {
	var encodeDecode int8
	var fileName string
	var fileContent []byte
	var fileReadError error
	fmt.Println("choose one")
	fmt.Println("1. encode")
	fmt.Println("2. decode")
	fmt.Scan(&encodeDecode)
	if encodeDecode != 1 && encodeDecode != 2 {
		log.Fatalln("please enter 1/2 only")
	}
	switch encodeDecode {
	case 1:
    /// encoding a file
		fmt.Printf("Enter filename to encode : ")
		fmt.Scan(&fileName)
		fileContent, fileReadError = os.ReadFile(fileName)
		if fileReadError != nil {
			log.Fatalln("error reading file / file doesn't exist")
		}
		encodedFileContent := util.Encode(fileContent)
		fmt.Println("encoding done!")
		fmt.Printf("Enter filename to store result : ")
		var outputFileName string
		fmt.Scan(&outputFileName)
		writeErr := os.WriteFile(outputFileName, encodedFileContent, 0666)
		if writeErr != nil {
			log.Fatalln("Error writting to file")
		}
		fmt.Println("file compressed successfully")
		break
	case 2:
    /// decoding a encoded file
		fmt.Printf("Enter filename to decode : ")
		fmt.Scan(&fileName)
		fileContent, fileReadError = os.ReadFile(fileName)
		if fileReadError != nil {
			log.Fatalln("error reading file / file doesn't exist")
		}
		var decodedString string = util.Decode(fileContent)
		fmt.Println("decoding done!")
		fmt.Printf("Enter filename to store result : ")
		var outputFileName string
		fmt.Scan(&outputFileName)
		writeErr := os.WriteFile(outputFileName, []byte(decodedString), 0666)
		if writeErr != nil {
			log.Fatalln("Error writting to file")
		}
		fmt.Println("file compressed successfully")
		break
	}
}
