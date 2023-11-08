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

	// t.Run("Cookware", func(t *testing.T) {
	// 	want := []Cookware{
	// 		{
	// 			Name:     "blender",
	// 			Quantity: "",
	// 		},
	// 	}
	// 	if !reflect.DeepEqual(res.Cookware, want) {
	// 		t.Errorf("want: %v\ngot: %v", want, res.Metadata)
	// 	}
	// })
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

func TestParseCookware(t *testing.T) {
	t.Run("NoCookware", func(t *testing.T) {
		got := parseCookware("Taste and adjust if desired.")

		if len(got) > 0 {
			t.Errorf("want: []\ngot: %v", got)
		}
	})

	t.Run("HasCookware", func(t *testing.T) {
		t.Run("CookwareAtEndNoPeriod", func(t *testing.T) {
			got := parseCookware("Place in #blender")
			want := []Cookware{
				{
					Name: "blender",
				},
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want: %v\ngot: %v", want, got)
			}
		})

		t.Run("CookwareAtEndWithPeriod", func(t *testing.T) {
			got := parseCookware("Place in #blender.")
			want := []Cookware{
				{
					Name: "blender",
				},
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want: %v\ngot: %v", want, got)
			}
		})

		t.Run("CookwareMidSentence", func(t *testing.T) {
			got := parseCookware("Place in #blender on high speed")
			want := []Cookware{
				{
					Name: "blender",
				},
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want: %v\ngot: %v", want, got)
			}
		})

		t.Run("CookwareWithEmptyQuantity", func(t *testing.T) {
			got := parseCookware("Place in #blender{} on high speed")
			want := []Cookware{
				{
					Name: "blender",
				},
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want: %v\ngot: %v", want, got)
			}
		})

		t.Run("CookwareWithMultiWordName", func(t *testing.T) {
			got := parseCookware("Place in #potato ricer{}")
			want := []Cookware{
				{
					Name: "potato ricer",
				},
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want: %v\ngot: %v", want, got)
			}
		})

		t.Run("MultipleCookware", func(t *testing.T) {
			got := parseCookware("Place in #potato ricer{} or #food processor{}")
			want := []Cookware{
				{
					Name: "potato ricer",
				},
				{
					Name: "food processor",
				},
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want: %v\ngot: %v", want, got)
			}
		})
	})
}
