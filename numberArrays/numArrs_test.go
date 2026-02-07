package numberarrays

import (
	"testing"
	"strings"
)


func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v want true", got)
	}
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		shay = Account{Name: "Shay", Balance: 100}
		justin = Account{Name: "Justin", Balance: 75}
		dave = Account{Name: "Dave", Balance: 200}

		transactions = []Transaction{
			NewTransaction(justin, shay, 100),
			NewTransaction(dave, justin, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(shay), 200.0)
	AssertEqual(t, newBalanceFor(justin), 0.0)
	AssertEqual(t, newBalanceFor(dave), 175.0)
	
}

func TestFind(t *testing.T) {
	type Person struct {
	Name	string
	}

	t.Run(" first first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x % 2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	t.Run("Find the Strongest Lifter", func(t *testing.T) {
		people := []Person{
			{Name: "Justin Davila"},
			{Name: "Darshay Blount"},
			{Name: "Joshua Blackwell"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Justin")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Justin Davila"})
	})
}
