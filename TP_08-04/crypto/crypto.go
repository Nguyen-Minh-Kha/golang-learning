package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"crypto/sha256"
	"strconv"
)

func compareHashes(hash1, hash2 []byte) bool {
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}
	return true
}

func main() {
	var tableHash [][]byte
	var index int
	for i:= 1; i < 4; i++ {
		content, err := ioutil.ReadFile("./image_" + strconv.Itoa(i) +".jpg")
		h := sha256.New()
		if err!= nil {
			log.Fatal(err)
		}
		h.Write([]byte(content))
		tableHash = append(tableHash, h.Sum(nil))
	}
	for i := len(tableHash)-1; i > 1 ; i-- {
		if !compareHashes(tableHash[i], tableHash[i-1]) {
			index = i
		}
	}
	fmt.Printf("image %v est unique", index)
}