package mlservice

import (
	"context"
	"fmt"
	"github.com/ernesto-jimenez/httplogger"
	"github.com/google/go-querystring/query"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type INITProClient interface {
	Auth(ctx context.Context, login, password string) (*AuthResponse, error)
	DocRegister(ctx context.Context, token, operation string, groupCode string, input interface{}) (*ResponseDocRegister, error)
	CheckOperationResult(ctx context.Context, token, uuid string, groupCode string) (*ResponseCheckOperationResult, error)
	GenerateFiscalCheck(uuid string) *string
}

type Options struct {
	Endpoint string
	Logger   *slog.Logger
}

func NewClient(opts *Options) INITProClient {

	return &clientImpl{
		opts: opts,
		client: &http.Client{
			Transport: httplogger.NewLoggedTransport(
				http.DefaultTransport,
				newLogger(opts.Logger)),
		},
	}

}

type clientImpl struct {
	opts     *Options
	token    string
	expireIn time.Time
	client   *http.Client
}

func (c *clientImpl) buildUri(p string, values interface{}) string {
	v, _ := query.Values(values)
	return strings.TrimRight(fmt.Sprintf(
		"%s/%s?%s",
		strings.TrimRight(c.opts.Endpoint, "/"),
		strings.TrimLeft(p, "/"),
		v.Encode(),
	), "?")
}
