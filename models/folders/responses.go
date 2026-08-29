package foldermodels

import sharedmodels "github.com/karman-digital/clickup/models/shared"

type FoldersResponse struct {
	Folders []Folder `json:"folders"`
}

type FolderTemplatesResponse struct {
	Templates []FolderTemplate `json:"templates"`
}

type FolderTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Folder struct {
	ID           sharedmodels.Scalar `json:"id"`
	Name         string              `json:"name"`
	Archived     bool                `json:"archived"`
	Hidden       bool                `json:"hidden"`
	ParentFolder sharedmodels.Scalar `json:"parent_folder"`
	Space        Space               `json:"space"`
	Lists        []FolderList        `json:"lists"`
}

type FolderList struct {
	ID        sharedmodels.Scalar `json:"id"`
	Name      string              `json:"name"`
	TaskCount int                 `json:"task_count"`
}

type Space struct {
	ID     sharedmodels.Scalar `json:"id"`
	Name   string              `json:"name"`
	Access bool                `json:"access"`
}
