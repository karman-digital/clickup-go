package listmodels

type ListCreationBody struct {
	Name                       string `json:"name"`
	Content                    string `json:"content,omitempty"`
	DueDate                    int64  `json:"due_date,omitempty"`
	DueDateTime                bool   `json:"due_date_time,omitempty"`
	Priority                   int    `json:"priority,omitempty"`
	Assignee                   int    `json:"assignee,omitempty"`
	Status                     string `json:"status,omitempty"`
	IncludeMarkdownDescription bool   `json:"include_markdown_description,omitempty"`
}
