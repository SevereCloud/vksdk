package object // import "github.com/severecloud/vksdk/object"

type ordersAmount struct {
	Amounts  []ordersAmountItem `json:"amounts"`
	Currency string             `json:"currency"`
}

type ordersAmountItem struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Votes       string `json:"votes"`
}

type ordersOrder struct {
	Amount              int    `json:"amount"`
	AppOrderID          int    `json:"app_order_id"`
	CancelTransactionID int    `json:"cancel_transaction_id"`
	Date                int    `json:"date"`
	ID                  int    `json:"id"`
	Item                string `json:"item"`
	ReceiverID          int    `json:"receiver_id"`
	Status              string `json:"status"`
	TransactionID       int    `json:"transaction_id"`
	UserID              int    `json:"user_id"`
}
