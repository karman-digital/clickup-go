package timetrackingmodels

import sharedmodels "github.com/karman-digital/clickup/models/shared"

type TimeTrackHistoryResponse struct {
	Data []TimeTrackHistory `json:"data"`
}

type TimeTrackHistory struct {
	ID     string `json:"id"`
	Field  string `json:"field"`
	Before string `json:"before"`
	After  string `json:"after"`
	Date   string `json:"date"`
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
}

type TimeEntryResponse struct {
	Data TimeEntryData `json:"data"`
}

type TimeEntriesResponse struct {
	Data []TimeEntryRecord `json:"data"`
}

type TimeEntryRecordResponse struct {
	Data TimeEntryRecord `json:"data"`
}

type TimeEntryRecord struct {
	ID           sharedmodels.Scalar `json:"id"`
	User         TimeEntryUser       `json:"user"`
	Start        sharedmodels.Scalar `json:"start"`
	End          sharedmodels.Scalar `json:"end"`
	Duration     sharedmodels.Scalar `json:"duration"`
	TaskLocation TimeEntryLocation   `json:"task_location"`
}

type TimeEntryUser struct {
	ID sharedmodels.Scalar `json:"id"`
}

type TimeEntryLocation struct {
	ListID  sharedmodels.Scalar `json:"list_id"`
	SpaceID sharedmodels.Scalar `json:"space_id"`
}

type TimeEntryData struct {
	ID           string   `json:"id"`
	Task         Task     `json:"task"`
	Wid          string   `json:"wid"`
	User         User     `json:"user"`
	Billable     bool     `json:"billable"`
	Start        int64    `json:"start"`
	End          string   `json:"end"`
	Duration     int64    `json:"duration"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	At           int64    `json:"at"`
	IsLocked     bool     `json:"is_locked"`
	TaskLocation any      `json:"task_location"`
}

type Task struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}

type TaskStatus struct {
	Status     string `json:"status"`
	ID         string `json:"id"`
	Color      string `json:"color"`
	Type       string `json:"type"`
	Orderindex int    `json:"orderindex"`
}

type User struct {
	ID             int     `json:"id"`
	Username       string  `json:"username"`
	Email          string  `json:"email"`
	Color          string  `json:"color"`
	Initials       string  `json:"initials"`
	ProfilePicture *string `json:"profilePicture"`
}
