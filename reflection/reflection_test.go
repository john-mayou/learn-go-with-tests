package reflection

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	tests := []struct {
		Name  string
		Input interface{}
		Calls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointer",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Paris"},
			},
			[]string{"London", "Paris"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Paris"},
			},
			[]string{"London", "Paris"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := []string{}

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			slices.Sort(got)
			slices.Sort(test.Calls)
			assert.Equal(t, got, test.Calls)
		})
	}

	t.Run("channels", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{33, "London"}
			channel <- Profile{34, "Paris"}
			close(channel)
		}()

		got := []string{}

		walk(channel, func(input string) {
			got = append(got, input)
		})

		slices.Sort(got)
		assert.Equal(t, got, []string{"London", "Paris"})
	})

	t.Run("function", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{33, "London"}, Profile{34, "Paris"}
		}

		got := []string{}

		walk(function, func(input string) {
			got = append(got, input)
		})

		slices.Sort(got)
		assert.Equal(t, got, []string{"London", "Paris"})
	})
}
