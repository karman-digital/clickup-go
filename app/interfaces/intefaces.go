package clickupintefaces

import (
	listmodels "github.com/karman-digital/clickup/models/lists"
	sharedmodels "github.com/karman-digital/clickup/models/shared"
	taskmodels "github.com/karman-digital/clickup/models/tasks"
	timetrackingmodels "github.com/karman-digital/clickup/models/timetracking"
)

type TimeEntry interface {
	GetTimeEntryHistory(id string) (timetrackingmodels.TimeTrackHistoryResponse, error)
	CreateTimeEntry(timeEntry timetrackingmodels.TimeEntry, opts ...sharedmodels.GetOptions) (timetrackingmodels.TimeEntryResponse, error)
	DeleteTimeEntry(id string) error
}

type Tasks interface {
	GetTask(taskID string, opts ...taskmodels.GetTaskOptions) (taskmodels.TaskGetResponse, error)
	CreateTask(listId string, task taskmodels.TaskPostBody) (taskmodels.TaskGetResponse, error)
}

type Lists interface {
	CreateList(folderId string, body listmodels.ListCreationBody) (listmodels.List, error)
	CreateFolderlessList(spaceId string, body listmodels.ListCreationBody) (listmodels.List, error)
	CreateFolderlessListFromTemplate(spaceId string, templateId string, body listmodels.ListCreationBody) (listmodels.List, error)
}
