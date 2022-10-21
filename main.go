package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const containerExt = "avi"

var changeMap = map[string]string{
	"DIVX": "FMP4",
	"XVID": "FMP4",
	"DIV5": "FMP4",
	"DIV3": "MP43",
}

func main() {
	inputPath := os.Args[1]

	files, err := listFiles(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		err = updateFourCC(file, changeMap)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func listFiles(path string) ([]string, error) {
	var files []string
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s: no such file or directory", path)
		}
		return nil, err
	}
	if fi.IsDir() {
		err := filepath.WalkDir(path, func(path string, dirEntry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !dirEntry.IsDir() {
				if matchFileType(dirEntry.Name()) {
					files = append(files, path)
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else if matchFileType(fi.Name()) {
		files = append(files, path)
	}

	return files, nil
}

func matchFileType(name string) bool {
	return strings.EqualFold(filepath.Ext(name), "."+containerExt)
}

func updateFourCC(path string, changeMap map[string]string) error {
	const fourccHeaderOffset = 112
	const fourccDescriptionOffset = 188

	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	buff := make([]byte, 4)
	_, err = file.ReadAt(buff, fourccDescriptionOffset)
	if err != nil {
		return fmt.Errorf("%s: can't read FourCC with error %s", path, err)
	}
	fourCCDescription := string(buff)
	if replace, found := changeMap[fourCCDescription]; found {
		fmt.Printf("FourCC %s replacing to %s...", fourCCDescription, replace)

		_, err = file.WriteAt([]byte(strings.ToLower(replace)), fourccHeaderOffset)
		if err != nil {
			fmt.Println("Error")
			return fmt.Errorf("%s: can't update FourCC with error %s", path, err)
		}
		_, err = file.WriteAt([]byte(replace), fourccDescriptionOffset)
		if err != nil {
			fmt.Println("Error")
			return fmt.Errorf("%s: can't update FourCC with error %s", path, err)
		}

		fmt.Println("OK")
	} else {
		fmt.Printf("FourCC %s skipped\n", fourCCDescription)
	}
	return nil
}
