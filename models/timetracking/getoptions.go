package timetrackingmodels

type GetOptions struct {
	IncludeTaskTags      bool `json:"include_task_tags"`
	IncludeLocationNames bool `json:"include_location_names"`
}
