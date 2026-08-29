package clickupintefaces

import (
	"context"

	foldermodels "github.com/karman-digital/clickup/models/folders"
	listmodels "github.com/karman-digital/clickup/models/lists"
	sharedmodels "github.com/karman-digital/clickup/models/shared"
	taskmodels "github.com/karman-digital/clickup/models/tasks"
	timetrackingmodels "github.com/karman-digital/clickup/models/timetracking"
)

type Folders interface {
	ListFolders(context.Context, string, bool) ([]foldermodels.Folder, error)
	GetFolder(context.Context, string) (foldermodels.Folder, error)
	MoveFolder(context.Context, string, string) error
	CreateFolderFromTemplate(context.Context, string, string, foldermodels.CreateFromTemplateBody) (foldermodels.Folder, error)
}

type TimeEntry interface {
	GetTimeEntryHistory(id string) (timetrackingmodels.TimeTrackHistoryResponse, error)
	GetTimeEntries(context.Context, timetrackingmodels.TimeEntryQuery) (timetrackingmodels.TimeEntriesResponse, error)
	GetTimeEntry(context.Context, string) (timetrackingmodels.TimeEntryRecordResponse, error)
	CreateTimeEntry(timeEntry timetrackingmodels.TimeEntry, opts ...sharedmodels.GetOptions) (timetrackingmodels.TimeEntryResponse, error)
	DeleteTimeEntry(id string) error
}

type Workspaces interface {
	GetWorkspaceUserIDs(context.Context) ([]string, error)
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
