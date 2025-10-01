package accounts

import (
	"fmt"
	"go-oo/holders"
)

type SavingsAccount struct {
	Holder                    holders.Holder
	Agency, Number, Operation int
	balance                   float64
}

func (a *SavingsAccount) Withdraw(withdrawValue float64) (string, bool) {
	if withdrawValue > a.balance {
		message := "Saldo insuficiente para realizar o saque de R$ " + fmt.Sprintf("%.2f", withdrawValue) + " da conta de " + a.Holder.Name
		return message, false
	}

	a.balance -= withdrawValue
	message := "Saque realizado com sucesso. Saldo atual da conta: R$ " + fmt.Sprintf("%.2f", a.balance) + " da conta de " + a.Holder.Name
	return message, true
}

func (a *SavingsAccount) Deposit(depositValue float64) (string, bool) {
	if depositValue <= 0 {
		message := "O valor de R$ " + fmt.Sprintf("%.2f", depositValue) + " é inválido para depósito na conta de " + a.Holder.Name
		return message, false
	}

	a.balance += depositValue
	message := "Depositando o valor de R$ " + fmt.Sprintf("%.2f", depositValue) + " na conta de " + a.Holder.Name
	return message, true
}

func (a *SavingsAccount) Transfer(transferValue float64, destinationAccount *Account) (string, bool) {
	withdrawMessage, withdrawStatus := a.Withdraw(transferValue)
	if withdrawStatus {
		return destinationAccount.Deposit(transferValue)
	}

	return withdrawMessage, false
}

func (a *SavingsAccount) GetBalance() float64 {
	return a.balance
}
