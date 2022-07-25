package accounts

import "errors"

type WellsFargo struct {
	balence int
}

func NewWellsFargo() *WellsFargo {
	return &WellsFargo{
		balence: 0,
	}
}

func (w *WellsFargo) GetBalance() int {
	return w.balence
}
func (w *WellsFargo) Deposit(amount int) {
	w.balence += amount
}
func (w *WellsFargo) Withdraw(amount int) error {
	newBalance := w.balence - amount
	if newBalance < 0 {
		return errors.New("Insufficient funds!")
	}

	w.balence = newBalance
	return nil
}