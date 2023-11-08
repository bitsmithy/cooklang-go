package main

import (
	"bufio"
	"bytes"
	"maps"
	"os"
	"strings"
)

const (
	metadataPrefix = ">> "
)

type Recipe struct {
	Metadata map[string]string
}

func ParseFile(path string) (Recipe, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Recipe{}, err
	}

	return Parse(b)
}

func Parse(b []byte) (Recipe, error) {
	r := Recipe{}
	m := map[string]string{}

	s := bufio.NewScanner(bytes.NewReader(b))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		maps.Copy(m, parseMetadata(line))
	}

	r.Metadata = m
	return r, nil
}

// Recognizes lines that look like:
// >> source: https://www.gimmesomeoven.com/baked-potato/
// >> time required: 1.5 hours
// >> course: dinner
// and adds it to the passed in map like:
//
//	{
//	  "source": "https://www.gimmesomeoven.com/baked-potato/",
//	  "time required": "1.5 hours",
//	  "course": "dinner"
//	}
func parseMetadata(line string) map[string]string {
	m := map[string]string{}

	if strings.HasPrefix(line, metadataPrefix) {
		line := strings.TrimPrefix(line, metadataPrefix)
		splits := strings.Split(line, ": ")
		key := splits[0]
		value := splits[1]
		m[key] = value
	}

	return m
}
