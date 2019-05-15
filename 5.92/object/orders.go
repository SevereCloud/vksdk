package object // import "github.com/SevereCloud/vksdk/5.92/object"

// OrdersAmount struct
type OrdersAmount struct {
	Amounts  []ordersAmountItem `json:"amounts"`
	Currency string             `json:"currency"` // Currency name
}

type ordersAmountItem struct {
	Amount      int    `json:"amount"`      // Votes amount in user's currency
	Description string `json:"description"` // Amount description
	Votes       string `json:"votes"`       // Votes number
}

// OrdersOrder struct
type OrdersOrder struct {
	Amount              int    `json:"amount"`                // Amount
	AppOrderID          int    `json:"app_order_id"`          // App order ID
	CancelTransactionID int    `json:"cancel_transaction_id"` // Cancel transaction ID
	Date                int    `json:"date"`                  // Date of creation in Unixtime
	ID                  int    `json:"id"`                    // Order ID
	Item                string `json:"item"`                  // Order item
	ReceiverID          int    `json:"receiver_id"`           // Receiver ID
	Status              string `json:"status"`                // Order status
	TransactionID       int    `json:"transaction_id"`        // Transaction ID
	UserID              int    `json:"user_id"`               // User ID
}
