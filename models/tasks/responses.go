package taskmodels

type TaskGetResponse struct {
	ID                  string        `json:"id,omitempty"`
	CustomID            string        `json:"custom_id,omitempty"`
	CustomItemID        int           `json:"custom_item_id,omitempty"`
	Name                string        `json:"name,omitempty"`
	TextContent         string        `json:"text_content,omitempty"`
	Description         string        `json:"description,omitempty"`
	Status              Status        `json:"status,omitempty"`
	OrderIndex          string        `json:"orderindex,omitempty"`
	DateCreated         string        `json:"date_created,omitempty"`
	DateUpdated         string        `json:"date_updated,omitempty"`
	DateClosed          string        `json:"date_closed,omitempty"`
	Creator             User          `json:"creator,omitempty"`
	Watchers            []User        `json:"watchers,omitempty"`
	GroupAssignees      []User        `json:"group_assignees,omitempty"`
	Checklists          []Checklist   `json:"checklists,omitempty"`
	Tags                []Tag         `json:"tags,omitempty"`
	Parent              string        `json:"parent,omitempty"`
	Priority            Priority      `json:"priority,omitempty"`
	DueDate             string        `json:"due_date,omitempty"`
	StartDate           string        `json:"start_date,omitempty"`
	Points              int           `json:"points,omitempty"`
	TimeEstimate        int           `json:"time_estimate,omitempty"`
	TimeSpent           int           `json:"time_spent,omitempty"`
	CustomFields        []CustomField `json:"custom_fields,omitempty"`
	List                Entity        `json:"list,omitempty"`
	Folder              Entity        `json:"folder,omitempty"`
	Space               Entity        `json:"space,omitempty"`
	URL                 string        `json:"url,omitempty"`
	MarkdownDescription string        `json:"markdown_description,omitempty"`
	Dependencies        []Dependency  `json:"dependencies,omitempty"`
	LinkedTasks         []LinkedTask  `json:"linked_tasks,omitempty"`
	Attachments         []Attachment  `json:"attachments,omitempty"`
	Assignees           []User        `json:"assignees,omitempty"`
}

type Dependency struct {
	TaskID      string `json:"task_id"`
	DependsOn   string `json:"depends_on"`
	Type        int    `json:"type"`
	DateCreated string `json:"date_created"`
	UserId      string `json:"user_id"`
	WorkspaceId string `json:"workspace_id"`
	ChainId     string `json:"chain_id"`
}

type LinkedTask struct {
	TaskID      string `json:"task_id"`
	LinkId      string `json:"link_id"`
	DateCreated string `json:"date_created"`
	UserId      string `json:"user_id"`
	WorkspaceId string `json:"workspace_id"`
}

type Priority struct {
	Color      string `json:"color"`
	Id         string `json:"id"`
	OrderIndex string `json:"orderindex"`
	Priority   string `json:"priority"`
}

type Tag struct {
	Name    string `json:"name"`
	TagFg   string `json:"tag_fg"`
	TabBg   string `json:"tab_bg"`
	Creator int    `json:"creator"`
}

type Status struct {
	Status     string  `json:"status"`
	Color      string  `json:"color"`
	OrderIndex float64 `json:"orderindex"`
	Type       string  `json:"type"`
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Color          string `json:"color"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profilePicture"`
}

type CustomField struct {
	ID             string         `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	Type           string         `json:"type,omitempty"`
	TypeConfig     map[string]any `json:"type_config,omitempty"`
	DateCreated    string         `json:"date_created,omitempty"`
	HideFromGuests bool           `json:"hide_from_guests,omitempty"`
	Value          any            `json:"value,omitempty"`
	Required       bool           `json:"required,omitempty"`
}

type Entity struct {
	ID string `json:"id"`
}

type Attachment struct {
	ID               string     `json:"id"`
	Date             string     `json:"date"`
	Title            string     `json:"title"`
	Type             int        `json:"type"`
	Source           int        `json:"source"`
	Version          int        `json:"version"`
	Extension        string     `json:"extension"`
	ThumbnailSmall   *string    `json:"thumbnail_small"`
	ThumbnailMedium  *string    `json:"thumbnail_medium"`
	ThumbnailLarge   *string    `json:"thumbnail_large"`
	IsFolder         *bool      `json:"is_folder"`
	MimeType         string     `json:"mimetype"`
	Hidden           bool       `json:"hidden"`
	ParentID         string     `json:"parent_id"`
	Size             int        `json:"size"`
	TotalComments    int        `json:"total_comments"`
	ResolvedComments int        `json:"resolved_comments"`
	User             User       `json:"user"`
	Deleted          bool       `json:"deleted"`
	Orientation      *string    `json:"orientation"`
	URL              string     `json:"url"`
	EmailData        *EmailData `json:"email_data"`
	URLWithQuery     string     `json:"url_w_query"`
	URLWithHost      string     `json:"url_w_host"`
}

type Checklist struct {
	ID          string  `json:"id"`
	TaskID      string  `json:"task_id"`
	Name        string  `json:"name"`
	DateCreated string  `json:"date_created"`
	OrderIndex  float64 `json:"orderindex"`
	Creator     int     `json:"creator"`
	Resolved    int     `json:"resolved"`
	Unresolved  int     `json:"unresolved"`
	Items       []Item  `json:"items"`
}

type Item struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	OrderIndex    float64 `json:"orderindex"`
	Assignee      User    `json:"assignee"`
	GroupAssignee *string `json:"group_assignee"`
	Resolved      bool    `json:"resolved"`
	Parent        *string `json:"parent"`
	DateCreated   string  `json:"date_created"`
	Children      []Item  `json:"children"`
}

type EmailData struct {
	ID          string   `json:"id"`
	Msg         string   `json:"msg"`
	From        string   `json:"from"`
	Email       string   `json:"email"`
	Client      string   `json:"client"`
	Subject     string   `json:"subject"`
	Attachments []string `json:"attachments"`
}
