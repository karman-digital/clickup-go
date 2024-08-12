package timetracking

import (
	"github.com/karman-digital/clickup/api/credentials"
	"github.com/karman-digital/clickup/api/timetracking/timeentry"
)

func NewTimeTracking(credentials *credentials.Credentials) TimeTracking {
	return TimeTracking{
		TimeEntry: timeentry.NewTimeEntryService(credentials),
	}
}
