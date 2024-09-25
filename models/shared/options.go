package sharedmodels

type GetOptions struct {
	CustomTaskIds              bool `json:"custom_task_ids"`
	TeamId                     int  `json:"team_id"`
	IncludeSubTasks            bool `json:"include_sub_tasks"`
	IncludeMarkdownDescription bool `json:"include_markdown_description"`
	IncludeTaskTags            bool `json:"include_task_tags"`
	IncludeLocationNames       bool `json:"include_location_names"`
}
