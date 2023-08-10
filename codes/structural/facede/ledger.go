package facede

import "fmt"

type ledger struct{}

func (l *ledger) makeEntry(accountID string, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d", accountID, txnType, amount)
	return
}
