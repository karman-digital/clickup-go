package workspacemodels

import sharedmodels "github.com/karman-digital/clickup/models/shared"

type AuthorizedWorkspacesResponse struct {
	Teams []Workspace `json:"teams"`
}

type Workspace struct {
	ID      sharedmodels.Scalar `json:"id"`
	Members []Member            `json:"members"`
}

type Member struct {
	User User `json:"user"`
}

type User struct {
	ID sharedmodels.Scalar `json:"id"`
}
