package listmodels

type List struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	OrderIndex     int            `json:"orderindex,omitempty"`
	Content        string         `json:"content,omitempty"`
	Status         Status         `json:"status,omitempty"`
	Priority       *Priority      `json:"priority,omitempty"`
	Assignee       *Assignee      `json:"assignee,omitempty"`
	TaskCount      *int           `json:"task_count,omitempty"`
	DueDate        *string        `json:"due_date,omitempty"`
	DueDateTime    bool           `json:"due_date_time,omitempty"`
	StartDate      *string        `json:"start_date,omitempty"`
	StartDateTime  *string        `json:"start_date_time,omitempty"`
	Folder         Folder         `json:"folder,omitempty"`
	Space          Space          `json:"space,omitempty"`
	Statuses       []StatusDetail `json:"statuses,omitempty"`
	InboundAddress string         `json:"inbound_address,omitempty"`
}

type Status struct {
	Status    string `json:"status"`
	Color     string `json:"color,omitempty"`
	HideLabel bool   `json:"hide_label,omitempty"`
}

type Priority struct {
	Priority string `json:"priority,omitempty"`
	Color    string `json:"color,omitempty"`
}

type Assignee struct {
	ID             int    `json:"id,omitempty"`
	Color          string `json:"color,omitempty"`
	Username       string `json:"username,omitempty"`
	Initials       string `json:"initials,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
}

type Folder struct {
	ID     string `json:"id"`
	Name   string `json:"name,omitempty"`
	Hidden bool   `json:"hidden,omitempty"`
	Access bool   `json:"access,omitempty"`
}

type Space struct {
	ID     string `json:"id"`
	Name   string `json:"name,omitempty"`
	Access bool   `json:"access,omitempty"`
}

type StatusDetail struct {
	Status     string `json:"status"`
	OrderIndex int    `json:"orderindex,omitempty"`
	Color      string `json:"color,omitempty"`
	Type       string `json:"type,omitempty"`
}
