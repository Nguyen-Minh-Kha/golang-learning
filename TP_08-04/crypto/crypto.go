package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func compareHashes(hash1, hash2 []byte) bool {
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}
	return true
}

func getFilesHashes(imagesFlags []string) [][]byte {
	var tableHash [][]byte
	for i:= 0; i < len(imagesFlags); i++ {
		fmt.Println(imagesFlags[i])
		content, err := os.ReadFile("./" + imagesFlags[i] +".jpg")
		h := sha256.New()
		if err!= nil {
			log.Fatal(err)
		}
		h.Write([]byte(content))
		tableHash = append(tableHash, h.Sum(nil))
	}
	return tableHash
}

func getUniqueImage(tableHash [][]byte) int {
	var index int
	for i := len(tableHash)-1; i > 1 ; i-- {
		if !compareHashes(tableHash[i], tableHash[i-1]) {
			index = i
		}
	}
	return index
} 

func main() {

	imagesFlag := flag.String("images", "", "list of images we want to compare (separated by commas)")

	flag.Parse()

	imagesFlagArray := strings.Split(*imagesFlag, ",")
	
	index := getUniqueImage(getFilesHashes(imagesFlagArray))
	
	fmt.Printf("image %v est unique", index)
}