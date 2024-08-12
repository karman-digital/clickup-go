package taskdata

import "github.com/karman-digital/clickup/api/credentials"

func NewTaskDataService(creds *credentials.Credentials) *TaskDataService {
	return &TaskDataService{creds}
}
