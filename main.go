package main

import (
	"fmt"
	"go-oo/accounts"
	"math/rand"
	"os"
)

func main() {
	andreAccount := accounts.Account{
		Holder:  "André",
		Agency:  1001,
		Number:  123456,
		Balance: 125.50,
	}

	joaoAccount := accounts.Account{
		Holder:  "João",
		Agency:  1001,
		Number:  123457,
		Balance: 265.45,
	}

	for {
		showMenu()
		option := readOption()

		switch option {
		case 1:
			fmt.Println(andreAccount)
			printSeparator()

			maxDepositValue := 200.
			minDepositValue := -20.
			depositValue := minDepositValue + rand.Float64()*(maxDepositValue-minDepositValue)
			depositMessage, _ := andreAccount.Deposit(depositValue)

			fmt.Println(depositMessage)
			printSeparator()

			fmt.Println(andreAccount)
			printSeparator()
		case 2:
			fmt.Println(andreAccount)
			printSeparator()

			maxWithdrawValue := 200.
			minWithdrawValue := 0.
			withdrawValue := minWithdrawValue + rand.Float64()*(maxWithdrawValue-minWithdrawValue)
			withdrawMessage, _ := andreAccount.Withdraw(withdrawValue)

			fmt.Println(withdrawMessage)
			printSeparator()

			fmt.Println(andreAccount)
			printSeparator()
		case 3:
			fmt.Println("Selecione a conta de origem:")
			fmt.Println("1 - André")
			fmt.Println("2 - João")
			fmt.Println("0 - Voltar")

			sourceOption := readOption()

			var sourceAccount, destinationAccount *accounts.Account

			issetSourceAccount := false

			switch sourceOption {
			case 1:
				sourceAccount = &andreAccount
				destinationAccount = &joaoAccount
				issetSourceAccount = true
			case 2:
				sourceAccount = &joaoAccount
				destinationAccount = &andreAccount
				issetSourceAccount = true
			case 0:
				break
			}

			if !issetSourceAccount {
				break
			}

			maxTransferValue := 200.
			minTransferValue := 0.
			transferValue := minTransferValue + rand.Float64()*(maxTransferValue-minTransferValue)
			transferMessage, _ := sourceAccount.Transfer(transferValue, destinationAccount)

			fmt.Println(transferMessage)
			fmt.Println(sourceAccount)
			fmt.Println(destinationAccount)
			printSeparator()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}
	}

}

func showMenu() {
	printSeparator()
	fmt.Println("1 - Depositar")
	fmt.Println("2 - Sacar")
	fmt.Println("3 - Transferir")
	fmt.Println("0 - Sair")
}

func readOption() int {
	var option int
	fmt.Scan(&option)

	return option
}

func printSeparator() {
	fmt.Println("------------------------------------------------------------------------------------")
}
