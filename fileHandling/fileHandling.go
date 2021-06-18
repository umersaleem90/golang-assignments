package fileHandling

import (
	"fmt"
	"log"
	"os"
)

var fileName string = "file.txt"

func Run() {
	// if !doesFileExists() {
	// 	writeFile("this is random text")
	// }
	writeFile("this is random text")
	readFile()
}

func readFile() {
	file, error := os.ReadFile(fileName)
  if error != nil {
    log.Fatal("Error: ", error)
    return
  }
  fmt.Println("File Contents:", string(file))
}

func writeFile(input string) {
	  file, error := os.Create(fileName)
    if error != nil {
        log.Fatal("Error while opening file: ", error)
    }
		defer file.Close()
		_, error = file.WriteString(input)
		if error != nil  {
			log.Fatal("Error while writing file: ", error)
		}
}

func doesFileExists() bool {
	if _, err := os.Stat(fileName); err == nil {
		writeFile("")
		return true
	} else {
		return false
	}
}

