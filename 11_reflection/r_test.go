package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Maz"},
			[]string{"Maz"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Maz", "Here"},
			[]string{"Maz", "Here"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Maz", 420},
			[]string{"Maz"},
		},
		{
			"nested fields",
			Person{"Maz", Profile{420, "Here"}},
			[]string{"Maz", "Here"},
		},
		{
			"pointers to things",
			&Person{
				"Maz",
				Profile{33, "Here"},
			},
			[]string{"Maz", "Here"},
		},
		{
			"slices",
			[]Profile{
				{33, "Here"},
				{34, "There"},
			},
			[]string{"Here", "There"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Here"},
				{34, "There"},
			},
			[]string{"Here", "There"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Here"}
			aChannel <- Profile{34, "There"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Here", "There"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Here"}, Profile{34, "There"}
		}

		var got []string
		want := []string{"Here", "There"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
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
