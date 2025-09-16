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


}
