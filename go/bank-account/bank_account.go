package account

import "sync"

//Account represents a bank account with the amount inside, its open/closed status and a mutex for ensuring safe concurrent transactions.
type Account struct {
	sync.RWMutex
	amount int64
	closed bool
}

//Close will close an account, paying out any left over funds. Further transactions on this account will no longer be possible.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	payout = a.amount
	a.amount = 0
	a.closed = true
	return payout, true
}

//Balance returns the current amount available within an account.
func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()
	if a.closed {
		return 0, false
	}
	return a.amount, true
}

//Deposit will add an amount (or subtract if negative) from an overall balance, disallowing overdrawing.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	if a.amount+amount < 0 {
		return 0, false
	}
	a.amount += amount
	return a.amount, true
}

//Open returns a pointer to a new account with amount greater than or equal to zero.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{amount: initialDeposit}
}
