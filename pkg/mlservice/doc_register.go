package mlservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DocRegisterRequest struct {
	// ExternalID - Внешний идентификатор
	ExternalID string `json:"external_id"`
	// Receipt - Чек
	Receipt Receipt `json:"receipt"`
	// Service - Сервисные данные
	Services Service `json:"services"`
	// Timestamp - Время запроса
	Timestamp string `json:"timestamp"`
}

type DocRegisterCorrectionRequest struct {
	// ExternalID - Внешний идентификатор
	ExternalID string `json:"external_id"`
	// Correction - Чек коррекции
	Correction Correction `json:"correction"`
	// Service - Сервисные данные
	Services Service `json:"services"`
	// Timestamp - Время запроса
	Timestamp string `json:"timestamp"`
}

func (c *clientImpl) DocRegister(ctx context.Context, token, operation, groupCode string, input interface{}) (*ResponseDocRegister, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.buildUri(fmt.Sprintf("/v1/%s/%s", groupCode, operation), nil),
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	req.Header.Set("Token", token)

	resp, err := jsonReq[ResponseDocRegister](c.client, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
