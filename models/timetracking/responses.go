package timetrackingmodels

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
