func NewAuthHandler() http.Handler {
	...
	...
	authRouter.GET(p("/cards"), cards.GetCards)
	authRouter.GET(p("/cards/:token"), cards.GetCard)
	authRouter.POST(p("/cards"), cards.CreateCard)
	authRouter.POST(p("/cards/:token"), cards.UpdateCard)

	authRouter.GET(p("/transactions/:id"), transactions.GetTransaction)
	authRouter.POST(p("/transactions"), transactions.CreateTransaction)
	authRouter.GET(p("/transactions"), transactions.GetTransactions)

	authRouter.POST(p("/transactions/:id/refund"), transactions.CreateRefund)
	...
	...
	...
	return mw
}
