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

type TimeTrackingData struct {
	Description string `json:"description"`
	IntervalId  string `json:"interval_id"`
}

type TimeTrackHistory struct {
	SharedHistoryItem
	Before TimeTrackState `json:"before"`
	After  TimeTrackState `json:"after"`
}
