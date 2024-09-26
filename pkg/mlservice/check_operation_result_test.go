package mlservice

import (
	"context"
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestClientImpl_CheckOperationResultStatus(t *testing.T) {
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

	t.Run("successful check result", func(t *testing.T) {
		resp, err := c.CheckOperationResult(timedCtx, token.Token, "c6ff42ff99c84b6b9bfc97956026e64c", TestGroupCode)
		logResponse(resp, err)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "wait", resp.Status)
	})

}
