package numberarrays

import (
	"testing"
)

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
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
