package accounts

import "fmt"

type Account struct {
	Holder  string
	Agency  int
	Number  int
	Balance float64
}

func (a *Account) Withdraw(withdrawValue float64) (string, bool) {
	if withdrawValue > a.Balance {
		message := "Saldo insuficiente para realizar o saque de R$ " + fmt.Sprintf("%.2f", withdrawValue) + " da conta de " + a.Holder
		return message, false
	}

	a.Balance -= withdrawValue
	message := "Saque realizado com sucesso. Saldo atual da conta: R$ " + fmt.Sprintf("%.2f", a.Balance) + " da conta de " + a.Holder
	return message, true
}

func (a *Account) Deposit(depositValue float64) (string, bool) {
	if depositValue <= 0 {
		message := "O valor de R$ " + fmt.Sprintf("%.2f", depositValue) + " é inválido para depósito na conta de " + a.Holder
		return message, false
	}

	a.Balance += depositValue
	message := "Depositando o valor de R$ " + fmt.Sprintf("%.2f", depositValue) + " na conta de " + a.Holder
	return message, true
}

func (a *Account) Transfer(transferValue float64, destinationAccount *Account) (string, bool) {
	withdrawMessage, withdrawStatus := a.Withdraw(transferValue)
	if withdrawStatus {
		return destinationAccount.Deposit(transferValue)
	}

	return withdrawMessage, false
}
