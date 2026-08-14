package timetrackingmodels

type GetOptions struct {
	IncludeTaskTags      bool `json:"include_task_tags"`
	IncludeLocationNames bool `json:"include_location_names"`
}

type TimeEntryQuery struct {
	StartDate            int64
	EndDate              int64
	AssigneeIDs          []string
	SpaceID              string
	IncludeLocationNames bool
}
