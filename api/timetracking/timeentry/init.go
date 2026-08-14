package timeentry

import "github.com/karman-digital/clickup/api/credentials"

func NewTimeEntryService(creds *credentials.Credentials) *TimeEntryService {
	return &TimeEntryService{Credentials: creds, reader: creds}
}

func newTimeEntryService(reader timeEntryReader) *TimeEntryService {
	return &TimeEntryService{reader: reader}
}
