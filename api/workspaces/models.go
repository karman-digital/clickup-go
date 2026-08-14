package workspaces

import (
	"context"
	"net/http"
)

type workspaceRequester interface {
	SendRequestWithContext(context.Context, string, string, []byte) (*http.Response, error)
}

type Service struct {
	requester workspaceRequester
	teamID    string
}
