package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
  	"io"
  	"log"
  	"net/http"
  	"os"
  	"strings"
  	"time"
)

func main() {

  SetupLogger()

  imagesFlag := flag.String("images", "", "list of images we want to compare (separated with coma)")
  urlsFlag := flag.String("urls", "", "list of images we want to compare (separated with coma)")
  flag.Parse()

  imagesFlagArray := strings.Split(*imagesFlag, ",")
  urlsFlagArray := strings.Split(*urlsFlag, ",")

  start := time.Now()
  for _, url := range urlsFlagArray {
    resp, errUrl := http.Get(url)
    if errUrl != nil {
      log.Default().Println(errUrl.Error())
    }
    imageName := getFileNameFromUrl(url)
    file, fileErr := os.Create(imageName)
    if fileErr != nil {
      log.Default().Println(fileErr.Error())
    }
    defer file.Close()
    io.Copy(file, resp.Body)
    defer resp.Body.Close()
    imagesFlagArray = append(imagesFlagArray, imageName)

  }

  elapsedTime := time.Since(start)
  fmt.Printf("It took: %v", elapsedTime)

  ArrayOfByteArray := buildArrayOfBytehArray(len(imagesFlagArray), imagesFlagArray)

  duplicatePosition := findDuplicatePosition(ArrayOfByteArray)

  UniqueImageArray := findHashOfUniqueImages(duplicatePosition)

  fmt.Println(UniqueImageArray)
}

func getFileNameFromUrl(url string) string {
  firstUrlSplit := strings.Split(url, "?")
  secondUrlSplit := strings.Split(firstUrlSplit[0], "/")
  arraySize := len(secondUrlSplit)
  imageName := secondUrlSplit[arraySize-1]
  return imageName
}

func SetupLogger() {
  log.SetPrefix("INFO: ")
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
  log.SetOutput(os.Stderr)
}

func findHashOfUniqueImages(duplicatePosition map[string][]int) []string {
  UniqueImageArray := make([]string, len(duplicatePosition))
  for index, intArray := range duplicatePosition {
    if len(intArray) == 1 {
      UniqueImageArray = append(UniqueImageArray, index)
    }
  }
  return UniqueImageArray
}

func findDuplicatePosition(ArrayOfByteArray [][]byte) map[string][]int {
  duplicatePosition := make(map[string][]int)

  for index, byteArray := range ArrayOfByteArray {
    duplicatePosition[string(byteArray)] = append(duplicatePosition[string(byteArray)], index)
  }

  return duplicatePosition
}

func buildArrayOfBytehArray(arraySize int, imagesFlagArray []string) [][]byte {
  ArrayOfByteArray := make([][]byte, arraySize)

  for i := 0; i < len(imagesFlagArray); i++ {
    bArray, err := getSHA256File(imagesFlagArray[i])
    if err != nil {
      log.Fatal(err.Error())
    }
    ArrayOfByteArray[i] = bArray

  }
  return ArrayOfByteArray
}

func readFile(filePath string) ([]byte, error) {
  bFile, err := os.ReadFile(filePath)
  return bFile, err
}

func getSHA256File(filePath string) ([]byte, error) {
  bFile, err := readFile(filePath)
  var hash []byte

  if err != nil {
    fmt.Println(err.Error())
  } else {
    hash32 := sha256.Sum256(bFile)
    hash = hash32[:]
  }

  return hash, err

}