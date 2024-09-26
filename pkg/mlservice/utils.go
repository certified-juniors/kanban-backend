package mlservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ernesto-jimenez/httplogger"
	"github.com/pkg/errors"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type httpLogger struct {
	log *slog.Logger
}

func (l httpLogger) LogRequest(req *http.Request) {
	if req.Body != nil {
		body, _ := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewReader(body))

		l.log.Info("Request",
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
			slog.String("body", string(body)),
		)
	} else {
		l.log.Info("Request",
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
		)
	}
}

func (l httpLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	durationMs := duration.Milliseconds()
	if err != nil {
		l.log.Error("Response error", slog.String("error", err.Error()))
	} else {
		l.log.Info("Response",
			slog.String("method", req.Method),
			slog.Int("status", res.StatusCode),
			slog.Int64("durationMs", durationMs),
			slog.String("url", req.URL.String()),
		)
	}
}

func newLogger(log *slog.Logger) httplogger.HTTPLogger {
	return httpLogger{
		log: log,
	}
}

// jsonReq осуществляет запрос и возвращает размаршифрованный ответ или ошибку
func jsonReq[T any](client *http.Client, req *http.Request) (*T, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "InitProClient.Do")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	var s T
	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	// Проверяем на наличие ошибок в JSON-ответе
	var respErr ResponseErr
	if err := json.Unmarshal(payload, &respErr); err == nil && respErr.ErrorID != "" {
		var errMsg string
		errMsg += fmt.Sprintf("[%d] %s ", respErr.Code, respErr.Text)

		return nil, fmt.Errorf("API error: %s", errMsg)
	}

	// Если не удалось размаршифровать ответ в структуру T, возвращаем оригинальный payload
	if err := json.Unmarshal(payload, &s); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s", payload)
	}

	return &s, nil
}
