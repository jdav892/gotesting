package numberarrays

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name	string
	Balance	float64
}

func Reduce[T, C any](collection []T, f func(C, T) C, initialValue C) C {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
			transactions,
			applyTransaction,
			account,
		)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}
