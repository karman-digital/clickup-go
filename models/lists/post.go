package listmodels

type ListCreationBody struct {
	Name                       string `json:"name"`
	Content                    string `json:"content"`
	DueDate                    int64  `json:"due_date"`
	DueDateTime                bool   `json:"due_date_time"`
	Priority                   int    `json:"priority"`
	Assignee                   int    `json:"assignee"`
	Status                     string `json:"status"`
	IncludeMarkdownDescription bool   `json:"include_markdown_description"`
}
