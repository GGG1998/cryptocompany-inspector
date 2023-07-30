package coingecko

import (
	"cryptocompany-inspector/pkg/website"
	"sync"
	"time"
)

type Coingecko struct {
	limitConnection int
	timeLimit       time.Duration
	job             chan website.Job
	quit            chan struct{}
	start           sync.Once
}
