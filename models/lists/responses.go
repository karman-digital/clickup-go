package listmodels

type List struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	OrderIndex     int            `json:"orderindex"`
	Content        string         `json:"content"`
	Status         Status         `json:"status"`
	Priority       Priority       `json:"priority"`
	Assignee       Assignee       `json:"assignee"`
	TaskCount      *int           `json:"task_count"`
	DueDate        string         `json:"due_date"`
	DueDateTime    bool           `json:"due_date_time"`
	StartDate      *string        `json:"start_date"`
	StartDateTime  *string        `json:"start_date_time"`
	Folder         Folder         `json:"folder"`
	Space          Space          `json:"space"`
	Statuses       []StatusDetail `json:"statuses"`
	InboundAddress string         `json:"inbound_address"`
}

type Status struct {
	Status    string `json:"status"`
	Color     string `json:"color"`
	HideLabel bool   `json:"hide_label"`
}

type Priority struct {
	Priority string `json:"priority"`
	Color    string `json:"color"`
}

type Assignee struct {
	ID             int    `json:"id"`
	Color          string `json:"color"`
	Username       string `json:"username"`
	Initials       string `json:"initials"`
	ProfilePicture string `json:"profilePicture"`
}

type Folder struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	Access bool   `json:"access"`
}

type Space struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Access bool   `json:"access"`
}

type StatusDetail struct {
	Status     string `json:"status"`
	OrderIndex int    `json:"orderindex"`
	Color      string `json:"color"`
	Type       string `json:"type"`
}
