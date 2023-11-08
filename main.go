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
	Cookware []Cookware
}

type Cookware struct {
	Name     string
	Quantity string
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
	c := []Cookware{}

	s := bufio.NewScanner(bytes.NewReader(b))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		maps.Copy(m, parseMetadata(line))
		c = append(c, parseCookware(line)...)
	}

	r.Metadata = m
	r.Cookware = c
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

// Recognizes lines that look like:
// Place the potatoes into a #pot.
// Mash the potatoes with a #potato masher{}.
// and extracts the cookware out into the passed in list:
//
//	[]Cookware{
//	  {
//	     Name: "pot",
//	     Quantity: "",
//	  },
//	  {
//	     Name: "potato masher",
//	     Quantity: "",
//	  },
//	}
func parseCookware(line string) []Cookware {
	startDelimiter := "#"
	explicitEndDelimiter := "{"
	implicitEndDelimiters := " ."
	startIndex := strings.Index(line, startDelimiter)

	if startIndex == -1 {
		return []Cookware{}
	}

	line = line[startIndex+1:]
	endIndex := strings.Index(line, explicitEndDelimiter)
	if endIndex == -1 {
		endIndex := strings.IndexAny(line, implicitEndDelimiters)
		if endIndex == -1 {
			return []Cookware{
				{Name: line},
			}
		}

		return append([]Cookware{
			{Name: line[:endIndex]},
		}, parseCookware(line[endIndex+1:])...)
	}

	return append([]Cookware{
		{Name: line[:endIndex]},
	}, parseCookware(line[endIndex+1:])...)
}
