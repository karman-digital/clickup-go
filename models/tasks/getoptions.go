package taskmodels

type GetTaskOptions struct {
	CustomTaskIds              bool `json:"custom_task_ids"`
	TeamId                     int  `json:"team_id"`
	IncludeSubTasks            bool `json:"include_sub_tasks"`
	IncludeMarkdownDescription bool `json:"include_markdown_description"`
}
