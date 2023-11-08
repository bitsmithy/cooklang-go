package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	f := "testdata/berry_smoothie.cook"
	res, err := ParseFile(f)
	if err != nil {
		t.Errorf("Could not parse file %s: %v", f, err)
	}

	t.Run("Metadata", func(t *testing.T) {
		want := map[string]string{
			"source":     "https://www.dinneratthezoo.com/wprm_print/6796",
			"total time": "6 minutes",
			"servings":   "2",
		}
		if !reflect.DeepEqual(res.Metadata, want) {
			t.Errorf("want: %v\ngot: %v", want, res.Metadata)
		}
	})
}

func TestParseMetadata(t *testing.T) {
	t.Run("IsMetadata", func(t *testing.T) {
		got := parseMetadata(">> course: dinner")
		want := map[string]string{
			"course": "dinner",
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v\ngot: %v", want, got)
		}
	})

	t.Run("NotMetadata", func(t *testing.T) {
		got := parseMetadata("Taste and add @honey{} if desired.")
		want := map[string]string{}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v\ngot: %v", want, got)
		}
	})
}
