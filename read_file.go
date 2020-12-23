package file_reader

import (
	"bufio"
    "log"
    "os"
)

func ReadLinebyLine() {
    file, err := os.Open("data_files/base_teste.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}