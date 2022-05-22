package main

import (
	"bytes"
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	filePath string
)

func main() {
	flag.StringVar(&filePath, "filepath", "", "Absolute path for file")

	flag.Parse()
	if len(filePath) < 1 {
		log.Fatal("No --filepath is given")
	}

	buf := new(bytes.Buffer)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Cannot open file")
	}
	defer file.Close()

	fstat, err := file.Stat()
	if err != nil {
		log.Fatal("Cannot get file information")
	}

	_, err = buf.ReadFrom(file)
	if err != nil {
		log.Fatal("Cannot read file")
	}
	fmt.Printf("%x\n", sha1.Sum([]byte(fmt.Sprintf("blob %d\x00%s", fstat.Size(), buf.String()))))
}
