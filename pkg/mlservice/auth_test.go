package mlservice

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_Auth(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 35*time.Second)
	defer cancel()

	c := CreateTestClient()

	resp, err := c.Auth(timedCtx, TestLogin, TestPassword)

	require.NoError(t, err)
	require.NotNil(t, resp)
}
