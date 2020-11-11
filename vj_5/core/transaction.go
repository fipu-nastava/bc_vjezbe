package core

import (
	. "../common"
	"fmt"
)

type Transaction struct {
	From Address
	To Address
	Amount int
	Signature []byte
}

func NewTransaction(from Address, to Address, amount int) (t *Transaction)  {
	return &Transaction{
		From: from,
		To: to,
		Amount:amount,
	}
}
func (t *Transaction) TxHash() Hash {
	return ValuesHash(
		t.From[:],
		t.To[:],
		[]byte(fmt.Sprintf("%d", t.Amount)[:]))
}
