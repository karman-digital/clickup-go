package lists

import "github.com/karman-digital/clickup/api/credentials"

func NewListService(creds *credentials.Credentials) *ListService {
	return &ListService{creds}
}
