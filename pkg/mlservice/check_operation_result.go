package mlservice

import (
	"context"
	"fmt"
	"net/http"
)

func (c *clientImpl) CheckOperationResult(ctx context.Context, token, uuid, groupCode string) (*ResponseCheckOperationResult, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri(fmt.Sprintf("/v1/%s/report/%s", groupCode, uuid), nil),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	req.Header.Set("Token", token)

	resp, err := jsonReq[ResponseCheckOperationResult](c.client, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
