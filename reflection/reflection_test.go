package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	
	cases := []struct {
		Name	string
		Input	interface{}
		ExpectedCalls	[]string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Justin"},
			[]string{"Justin"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Justin", "Orlando"},
			[]string{"Justin", "Orlando"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age int
			}{"Justin", 31},
			[]string{"Justin"},
		},
		{
			"nested fields",
			Person {
				"Justin",
				Profile{31, "Orlando"},
			},
			[]string{"Justin", "Orlando"},
		},
		{
			"pointers to things",
			&Person{
				"Justin",
				Profile{31, "Orlando"},
			},
			[]string{"Justin", "Orlando"},
		},
		{
			"slices",
			[]Profile {
				{31, "Orlando"},
				{32, "Deltona"},
			},
			[]string{"Orlando", "Deltona"},
		},
		{
			"arrays",
			[2]Profile {
				{31, "Orlando"},
				{32, "Deltona"},
			},
			[]string{"Orlando", "Deltona"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t * testing.T) {
		aMap := map[string]string{
			"Cow": "Moo",
			"Goat": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}

}
