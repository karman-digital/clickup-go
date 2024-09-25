package timetrackingmodels

type TimeEntry struct {
	Description string `json:"description"`
	Tags        []struct {
		Name  string `json:"name"`
		TagBg string `json:"tag_bg"`
		TagFg string `json:"tag_fg"`
	} `json:"tags,omitempty"`
	Start    int64  `json:"start"`
	Billable bool   `json:"billable"`
	Duration int64  `json:"duration"`
	Assignee int    `json:"assignee"`
	TaskID   string `json:"tid"`
	End      int64  `json:"end,omitempty"`
}
