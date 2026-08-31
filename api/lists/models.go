package lists

import (
	"net/http"

	"github.com/karman-digital/clickup/api/credentials"
)

type ListService struct {
	*credentials.Credentials
	send func(string, string, []byte) (*http.Response, error)
}
