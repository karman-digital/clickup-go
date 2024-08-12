package tasks

import (
	"github.com/karman-digital/clickup/api/credentials"
	"github.com/karman-digital/clickup/api/tasks/taskdata"
)

func NewTasks(creds *credentials.Credentials) *Tasks {
	return &Tasks{
		taskdata.NewTaskDataService(creds),
	}
}
