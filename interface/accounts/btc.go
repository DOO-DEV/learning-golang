package accounts

import "errors"

type Btc struct {
	balence int
	fee int
}

func NewBtcAccount() *Btc {
	return &Btc{
		balence: 0,
		fee: 300,
	}
}

func (b *Btc) GetBalance() int {
	return b.balence
}
func (b *Btc) Deposit(amount int) {
	b.balence += amount
}
func (b *Btc) Withdraw(amount int) error {
	newBalance := b.balence - amount - b.fee
	if newBalance < 0 {
		return errors.New("Insufficient funds!")
	}

	b.balence = newBalance
	return nil
}