package payments_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/payments"
	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	t.Parallel()

	f := func(err payments.Error, want string) {
		t.Helper()

		assert.EqualError(t, err, want)
	}

	f(
		payments.Error{
			Code: payments.CommonError,
			Msg:  "text",
		},
		"payments: text",
	)
	f(
		payments.Error{
			Code: payments.CommonError,
		},
		"payments: error with code 1",
	)
}

func TestOrderStatusChangeRequest_OfferID(t *testing.T) {
	t.Parallel()

	f := func(r payments.OrderStatusChangeRequest, wantOfferID int) {
		offerID := r.OfferID()
		assert.Equal(t, wantOfferID, offerID)
	}

	f(
		payments.OrderStatusChangeRequest{
			Item: "test1234",
		},
		0,
	)
	f(
		payments.OrderStatusChangeRequest{
			Item: "offer_1234",
		},
		1234,
	)
	f(
		payments.OrderStatusChangeRequest{
			Item: "offer_1234a",
		},
		0,
	)
}
