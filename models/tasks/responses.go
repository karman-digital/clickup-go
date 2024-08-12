package taskmodels

type Task struct {
	ID                  string        `json:"id"`
	CustomID            string        `json:"custom_id"`
	CustomItemID        int           `json:"custom_item_id"`
	Name                string        `json:"name"`
	TextContent         string        `json:"text_content"`
	Description         string        `json:"description"`
	Status              Status        `json:"status"`
	OrderIndex          string        `json:"orderindex"`
	DateCreated         string        `json:"date_created"`
	DateUpdated         string        `json:"date_updated"`
	DateClosed          string        `json:"date_closed"`
	Creator             Creator       `json:"creator"`
	Assignees           []string      `json:"assignees"`
	Watchers            []string      `json:"watchers"`
	Checklists          []string      `json:"checklists"`
	Tags                []string      `json:"tags"`
	Parent              string        `json:"parent"`
	Priority            string        `json:"priority"`
	DueDate             string        `json:"due_date"`
	StartDate           string        `json:"start_date"`
	Points              int           `json:"points"`
	TimeEstimate        string        `json:"time_estimate"`
	TimeSpent           string        `json:"time_spent"`
	CustomFields        []CustomField `json:"custom_fields"`
	List                Entity        `json:"list"`
	Folder              Entity        `json:"folder"`
	Space               Entity        `json:"space"`
	URL                 string        `json:"url"`
	MarkdownDescription string        `json:"markdown_description"`
	Attachments         []Attachment  `json:"attachments"`
}

type Status struct {
	Status     string `json:"status"`
	Color      string `json:"color"`
	OrderIndex int    `json:"orderindex"`
	Type       string `json:"type"`
}

type Creator struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profilePicture"`
}

type CustomField struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	TypeConfig     map[string]any `json:"type_config"`
	DateCreated    string         `json:"date_created"`
	HideFromGuests bool           `json:"hide_from_guests"`
	Value          Value          `json:"value"`
	Required       bool           `json:"required"`
}

type Value struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	Initials       string `json:"initials"`
	ProfilePicture string `json:"profilePicture"`
}

type Entity struct {
	ID string `json:"id"`
}

type Attachment struct {
	ID               string       `json:"id"`
	Date             int64        `json:"date"`
	Title            string       `json:"title"`
	Type             int          `json:"type"`
	Source           int          `json:"source"`
	Version          int          `json:"version"`
	Extension        string       `json:"extension"`
	ThumbnailSmall   *string      `json:"thumbnail_small"`
	ThumbnailMedium  *string      `json:"thumbnail_medium"`
	ThumbnailLarge   *string      `json:"thumbnail_large"`
	IsFolder         *bool        `json:"is_folder"`
	MimeType         string       `json:"mimetype"`
	Hidden           bool         `json:"hidden"`
	ParentID         string       `json:"parent_id"`
	Size             int          `json:"size"`
	TotalComments    int          `json:"total_comments"`
	ResolvedComments int          `json:"resolved_comments"`
	User             []UserDetail `json:"user"`
	Deleted          bool         `json:"deleted"`
	Orientation      *int         `json:"orientation"`
	URL              string       `json:"url"`
	EmailData        *string      `json:"email_data"`
	URLWithQuery     string       `json:"url_w_query"`
	URLWithHost      string       `json:"url_w_host"`
}

type UserDetail struct {
	ID             []IDDetail `json:"id"`
	Username       string     `json:"username"`
	Initials       string     `json:"initials"`
	Email          string     `json:"email"`
	Color          string     `json:"color"`
	ProfilePicture string     `json:"profilePicture"`
}

type IDDetail struct {
	Type            string `json:"type"`
	ContentEncoding string `json:"contentEncoding"`
}
