package mlservice

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestClientImpl_DocRegisterStatus(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	c := CreateTestClient()

	logResponse := func(resp interface{}, err error) {
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			respJSON, _ := json.MarshalIndent(resp, "", "  ")
			log.Printf("Response: %s", respJSON)
		}
	}

	token, err := c.Auth(timedCtx, TestLogin, TestPassword)
	if err != nil {
		logResponse(token, err)
	}

	newReceipt := func() *DocRegisterRequest {
		return &DocRegisterRequest{
			ExternalID: uuid.New().String(),
			Receipt: Receipt{
				Client: Client{
					Phone: "+79876543210",
				},
				Company: Company{
					Email:          "chek@romashka.ru",
					Sno:            "osn",
					Inn:            "5902034504",
					PaymentAddress: "http://magazin.ru/",
				},
				Items: []Item{
					{
						Name:     "колбаса Клинский Брауншвейгская с/к в/с ",
						Price:    1000.00,
						Quantity: 1,
						Sum:      300.00,
						Vat: Vat{
							Type: "vat20",
						},
					},
					{
						Name:     "колбаса Клинский Брауншвейгская с/к в/с ",
						Price:    100.00,
						Quantity: 1.0,
						Sum:      100.00,
						Vat: Vat{
							Type: "vat20",
						},
					},
				},
				Payments: []Payment{
					{
						Type: 1, // БЕЗНАЛИЧНЫЙ
						Sum:  400.0,
					},
				},
				Total: 400.0,
			},
			Services: Service{
				CallbackUrl: "https://example.com/callback",
			},
			Timestamp: "01.02.17 13:45:00",
		}
	}

	newCorrection := func(correctionType string) *DocRegisterCorrectionRequest {
		return &DocRegisterCorrectionRequest{
			ExternalID: uuid.New().String(),
			Correction: Correction{
				Company: Company{
					Email:          "sdf",
					Sno:            "osn",
					Inn:            "5902034504",
					PaymentAddress: "http://magazin.ru/",
				},
				CorrectionInfo: CorrectionInfo{
					Type:       "correctionType",
					BaseDate:   "01.02.17",
					BaseNumber: "1",
					BaseName:   "name",
				},
				Payments: []Payment{
					{
						Type: 1, // БЕЗНАЛИЧНЫЙ
						Sum:  400.0,
					},
				},
				Vats: []Vat{
					{
						Type: "vat20",
						Sum:  80.0,
					},
				},
				Cashier: "cashier",
			},
			Services: Service{
				CallbackUrl: "https://example.com/callback",
			},
			Timestamp: "01.02.17 13:45:00",
		}
	}

	t.Run("successful sell request", func(t *testing.T) {
		receipt := newReceipt()
		resp, err := c.DocRegister(timedCtx, token.Token, "sell", TestGroupCode, receipt)
		logResponse(resp, err)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "wait", resp.Status)
	})

	t.Run("successful sell_refund request", func(t *testing.T) {
		receipt := newReceipt()
		resp, err := c.DocRegister(timedCtx, token.Token, "sell_refund", TestGroupCode, receipt)
		logResponse(resp, err)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "wait", resp.Status)
	})

	t.Run("successful correction_income request", func(t *testing.T) {
		correction := newCorrection("self")
		resp, err := c.DocRegister(timedCtx, token.Token, "correction_income", TestGroupCode, correction)
		logResponse(resp, err)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "wait", resp.Status)
	})

	t.Run("successful correction_expense request", func(t *testing.T) {
		correction := newCorrection("instruction")
		resp, err := c.DocRegister(timedCtx, token.Token, "correction_expense", TestGroupCode, correction)
		logResponse(resp, err)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "wait", resp.Status)
	})
}
