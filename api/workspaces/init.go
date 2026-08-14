package workspaces

import "github.com/karman-digital/clickup/api/credentials"

func NewWorkspaceService(credentials *credentials.Credentials) *Service {
	return newWorkspaceService(credentials, credentials.GetTeamId())
}

func newWorkspaceService(requester workspaceRequester, teamID string) *Service {
	return &Service{requester: requester, teamID: teamID}
}
