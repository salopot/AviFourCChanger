package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputPath := os.Args[1]
	fmt.Println(inputPath)

	var changeMap = map[string]string{
		"DIVX": "FMP4",
		"XVID": "FMP4",
		"DIV5": "FMP4",
		"DIV3": "MP43",
	}

	updateFourCC(inputPath, changeMap)
}

func updateFourCC(name string, changeMap map[string]string) {
	const fourccHeaderOffset = 112
	const fourccDescriptionOffset = 188

	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	buff := make([]byte, 4)
	file.ReadAt(buff, fourccDescriptionOffset)
	fourCCDescription := string(buff)

	if replace, found := changeMap[fourCCDescription]; found {
		fmt.Printf("Codec %s replcacing to %s...", fourCCDescription, replace)

		file.WriteAt([]byte(strings.ToLower(replace)), fourccHeaderOffset)
		file.WriteAt([]byte(replace), fourccDescriptionOffset)

		fmt.Println("OK")
	} else {
		fmt.Printf("Codec %s skipped\n", fourCCDescription)
	}
}
