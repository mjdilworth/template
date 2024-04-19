package apploader

import (
	"context"
)

type App interface {
	Run(context.Context)
}
