package website

import (
	"context"
	"cryptocompany-inspector/pkg/redis"
	"fmt"
	"io"
	"net/http"
)

type Job interface {
	Execute(context.Context) error
	OnFail()
}

type ApiCallJob struct {
	id  string
	url string

	// TODO: replace with interface
	httpClient *http.Client
	redis      *redis.Redis
}

func (a ApiCallJob) Execute(ctx context.Context) error {

	response, err := http.Get(a.url)
	if err != nil {
		return fmt.Errorf("ApiCallJob http: %w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("ApiCallJob http read response: %w", err)
	}

	return a.redis.LPush(ctx, a.id, string(body))
}
