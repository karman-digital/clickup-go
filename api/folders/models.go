package folders

import (
	"context"
	"net/http"
)

type requester interface {
	SendRequestWithContext(context.Context, string, string, []byte) (*http.Response, error)
}

type Service struct {
	requester requester
}
