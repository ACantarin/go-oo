package accounts

type BillPayer interface {
	Withdraw(float64) (string, bool)
}
