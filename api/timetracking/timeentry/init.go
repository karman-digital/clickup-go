package timeentry

import "github.com/karman-digital/clickup/api/credentials"

func NewTimeEntryService(creds *credentials.Credentials) *TimeEntryService {
	return &TimeEntryService{creds}
}
