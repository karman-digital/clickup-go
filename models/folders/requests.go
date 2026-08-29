package foldermodels

type MoveFolderRequest struct {
	SpaceID string `json:"space_id"`
}

type CreateFromTemplateBody struct {
	Name string `json:"name"`
}
