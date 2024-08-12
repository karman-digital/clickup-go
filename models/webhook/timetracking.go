package webhookmodels

type TimeTrackingPayload struct {
	CoreTaskPayloadData
	HistoryItems []TimeTrackHistory `json:"history_items"`
	Data         TimeTrackingData   `json:"data"`
}

type TimeTrackState struct {
	ID        string `json:"id"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Time      string `json:"time"`
	Source    string `json:"source"`
	DateAdded string `json:"date_added"`
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	Initials       string `json:"initials"`
	ProfilePicture string `json:"profilePicture"`
}

type TimeTrackingData struct {
	Description string `json:"description"`
	IntervalId  string `json:"interval_id"`
}

type TimeTrackHistory struct {
	SharedHistoryItem
	Start TimeTrackState `json:"start"`
	End   TimeTrackState `json:"end"`
}
