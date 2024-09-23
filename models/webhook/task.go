package webhookmodels

type TaskUpdatedPayload struct {
	CoreTaskPayloadData
	HistoryItems []TaskUpdatedHistory `json:"history_items"`
}

type TaskUpdatedHistory struct {
	SharedHistoryItem
	Before      interface{}  `json:"before"`
	After       interface{}  `json:"after"`
	CustomField *CustomField `json:"custom_field,omitempty"`
}

type CustomField struct {
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	Type               string      `json:"type"`
	TypeConfig         TypeConfig  `json:"type_config"`
	ValuesSet          interface{} `json:"values_set"`
	UserID             string      `json:"userid"`
	DateCreated        string      `json:"date_created"`
	HideFromGuests     bool        `json:"hide_from_guests"`
	TeamID             string      `json:"team_id"`
	Deleted            bool        `json:"deleted"`
	DeletedBy          interface{} `json:"deleted_by"`
	Pinned             bool        `json:"pinned"`
	Required           bool        `json:"required"`
	RequiredOnSubtasks bool        `json:"required_on_subtasks"`
	LinkedSubcategory  interface{} `json:"linked_subcategory"`
}

type TypeConfig struct {
	Default     int         `json:"default"`
	Placeholder interface{} `json:"placeholder"`
	NewDropDown bool        `json:"new_drop_down"`
	Options     []Option    `json:"options"`
}

type Option struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Value      string      `json:"value"`
	Type       string      `json:"type"`
	Color      interface{} `json:"color"`
	OrderIndex int         `json:"orderindex"`
}
