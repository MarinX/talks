// Processor is the interface for payment processors
type Processor interface {
	Name() string
	Charge(amount float64, currency string, description string,
		card *types.CardDB, cvv string, metadata map[string]string) *Result

	Auth(amount float64, currency string, description string,
		card *types.CardDB, cvv string, metadata map[string]string) *Result

	Refund(transaction *types.TransactionDB, amount float64) *Result
	Void() *Result
	GetTransactionByProcessorID(transactionID string) (*types.TransactionDB, error)
	GetTransactionsForCustomer(customerID string) ([]*types.TransactionDB, error)
}
