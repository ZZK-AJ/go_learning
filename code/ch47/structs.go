package profiling

type Request struct {
	TransactionID string `jsontest:"transaction_id"`
	PayLoad       []int  `jsontest:"payload"`
}

type Response struct {
	TransactionID string `jsontest:"transaction_id"`
	Expression    string `jsontest:"exp"`
}
