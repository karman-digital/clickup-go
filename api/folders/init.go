package folders

import "github.com/karman-digital/clickup/api/credentials"

func NewService(credentials *credentials.Credentials) *Service {
	return &Service{requester: credentials}
}
