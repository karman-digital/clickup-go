package taskmodels

type TaskPostBody struct {
	Name                      string        `json:"name"`
	Description               string        `json:"description,omitempty"`
	Assignees                 []int         `json:"assignees,omitempty"`
	Archived                  bool          `json:"archived"`
	GroupAssignees            []int         `json:"group_assignees,omitempty"`
	Tags                      []string      `json:"tags,omitempty"`
	Status                    string        `json:"status,omitempty"`
	Priority                  *int          `json:"priority,omitempty"`
	DueDate                   int64         `json:"due_date,omitempty"`
	DueDateTime               bool          `json:"due_date_time"`
	TimeEstimate              int32         `json:"time_estimate,omitempty"`
	StartDate                 int64         `json:"start_date,omitempty"`
	StartDateTime             bool          `json:"start_date_time"`
	Points                    float64       `json:"points,omitempty"`
	NotifyAll                 bool          `json:"notify_all"`
	Parent                    *string       `json:"parent,omitempty"`
	LinksTo                   *string       `json:"links_to,omitempty"`
	CheckRequiredCustomFields bool          `json:"check_required_custom_fields,omitempty"`
	CustomFields              []CustomField `json:"custom_fields,omitempty"`
	CustomItemID              *int          `json:"custom_item_id,omitempty"`
}
