package main

import (
	"fmt"
	"go-oo/accounts"
	"go-oo/holders"
	"math/rand"
	"os"
)

func main() {
	andreHolder := holders.Holder{
		Name:      "André",
		Document:  "123.456.789-00",
		Ocupation: "Programador",
	}
	andreAccount := accounts.Account{
		Holder:  andreHolder,
		Agency:  1001,
		Number:  123456,
		Balance: 125.50,
	}

	joaoHolder := holders.Holder{
		Name:      "João",
		Document:  "987.654.321-00",
		Ocupation: "Product Owner",
	}

	joaoAccount := accounts.Account{
		Holder:  joaoHolder,
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
		case 4:
			WithdrawFromRunningPayer()
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
	fmt.Println("4 - Pagar conta")
	fmt.Println("0 - Sair")
}

func readOption() int {
	var option int
	fmt.Scan(&option)

	return option
}

func WithdrawFromRunningPayer() {
	johnDoeAccount := accounts.Account{}
	johnDoeAccount.Deposit(1000)

	fmt.Println("Saldo da conta de John Doe antes de pagar a conta:", johnDoeAccount.GetBalance())
	payBill(&johnDoeAccount)
	fmt.Println("Saldo da conta de John Doe depois de pagar a conta:", johnDoeAccount.GetBalance())

	janeDoeAccount := accounts.SavingsAccount{}
	janeDoeAccount.Deposit(500)

	fmt.Println("Saldo da conta de Jane Doe antes de pagar a conta:", janeDoeAccount.GetBalance())
	payBill(&janeDoeAccount)
	fmt.Println("Saldo da conta de Jane Doe depois de pagar a conta:", janeDoeAccount.GetBalance())
}

func payBill(billPayer accounts.BillPayer) {
	billPayer.Withdraw(100)
}

func printSeparator() {
	fmt.Println("------------------------------------------------------------------------------------")
}
