package foldermodels

import sharedmodels "github.com/karman-digital/clickup/models/shared"

type FoldersResponse struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	ID           sharedmodels.Scalar `json:"id"`
	Name         string              `json:"name"`
	Archived     bool                `json:"archived"`
	Hidden       bool                `json:"hidden"`
	ParentFolder sharedmodels.Scalar `json:"parent_folder"`
	Space        Space               `json:"space"`
}

type Space struct {
	ID     sharedmodels.Scalar `json:"id"`
	Name   string              `json:"name"`
	Access bool                `json:"access"`
}
