package apploader

import (
	"context"
)

type testApp struct{}

func (t testApp) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
