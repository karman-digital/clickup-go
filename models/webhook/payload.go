package webhookmodels

type CoreTaskPayloadData struct {
	Event     string `json:"event"`
	TaskID    string `json:"task_id"`
	WebhookID string `json:"webhook_id"`
}

type SharedHistoryItem struct {
	ID       string `json:"id"`
	Type     int    `json:"type"`
	Date     string `json:"date"`
	Field    string `json:"field"`
	ParentID string `json:"parent_id"`
	Data     Data   `json:"data"`
	Source   string `json:"source"`
	User     User   `json:"user"`
}

type Data map[string]any
