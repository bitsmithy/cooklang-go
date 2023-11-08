package main

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

type Spec struct {
	Version int             `yaml:"version"`
	Tests   map[string]Test `yaml:"tests"`
}

type Test struct {
	Source string `yaml:"source"`
	Result Result `yaml:"result"`
}

type Result struct {
	Steps [][]struct {
		Type     string `yaml:"type"`
		Value    string `yaml:"value"`
		Quantity string `yaml:"quantity"`
		Name     string `yaml:"name"`
		Units    string `yaml:"units"`
	} `yaml:"steps"`
	Metadata map[string]string `yaml:"metadata"`
}

func TestCanonical(t *testing.T) {
	b, err := os.ReadFile("testdata/canonical.yaml")
	if err != nil {
		t.Fatalf("could not read canonical specs file: %v", err)
	}

	var spec Spec
	if err := yaml.Unmarshal(b, &spec); err != nil {
		t.Fatalf("could not parse canonical specs: %v", err)
	}

	for name, test := range spec.Tests {
		t.Run(name, func(t *testing.T) {
			_, err := Parse([]byte(test.Source))
			if err != nil {
				t.Errorf("could not parse test: %v", err)
			}

			// if got != test.Source {
			// 	t.Errorf("%s did not parse correctly", name)
			// }
		})
	}
}
