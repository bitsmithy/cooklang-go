package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	canonicalSource = "https://raw.githubusercontent.com/cooklang/spec/main/tests/canonical.yaml"
	testdataDir     = "testdata"
)

func main() {
	err := os.Mkdir(testdataDir, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("could not create testdata directory: %v", err)
	}

	resp, err := http.Get(canonicalSource)
	if err != nil {
		log.Fatalf("could not fetch canonical tests: %v", err)
	}
	defer resp.Body.Close()

	f, err := os.Create(fmt.Sprintf("%s/canonical.yaml", testdataDir))
	if err != nil {
		log.Fatalf("could not create canonical test file: %v", err)
	}
	defer f.Close()

	_, err = f.ReadFrom(resp.Body)
	if err != nil {
		log.Fatalf("could not write canonical test file: %v", err)
	}
}
