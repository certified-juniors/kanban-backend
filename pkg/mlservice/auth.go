package mlservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"net/http"
)

type AuthCredentials struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

type AuthResponse struct {
	Err   *ResponseErr `json:"error"`
	Token string       `json:"token"`
	Time  string       `json:"timestamp"`
}

func (c *clientImpl) Auth(ctx context.Context, login, password string) (*AuthResponse, error) {
	if len(login) == 0 || len(password) == 0 {
		return nil, fmt.Errorf("empty credentials")
	}

	authCredentials := AuthCredentials{
		Login:    login,
		Password: password,
	}

	payload, err := json.Marshal(authCredentials)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		c.buildUri("/v1/getToken", nil),
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return jsonReq[AuthResponse](c.client, req)
}
