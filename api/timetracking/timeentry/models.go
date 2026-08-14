package timeentry

import (
	"context"
	"net/http"

	"github.com/karman-digital/clickup/api/credentials"
)

type timeEntryReader interface {
	GetTeamId() string
	SendTimeTrackingRequestWithContext(context.Context, string, string, []byte) (*http.Response, error)
}

type TimeEntryService struct {
	*credentials.Credentials
	reader timeEntryReader
}
