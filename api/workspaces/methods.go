package workspaces

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	workspacemodels "github.com/karman-digital/clickup/models/workspaces"
)

func (service *Service) GetWorkspaceUserIDs(ctx context.Context) ([]string, error) {
	response, err := service.requester.SendRequestWithContext(ctx, http.MethodGet, "/team", nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error body: %s", string(body))
	}
	var workspaces workspacemodels.AuthorizedWorkspacesResponse
	if err := json.Unmarshal(body, &workspaces); err != nil {
		return nil, err
	}
	for _, workspace := range workspaces.Teams {
		if workspace.ID.String() != service.teamID {
			continue
		}
		users := make([]string, 0, len(workspace.Members))
		for _, member := range workspace.Members {
			if id := member.User.ID.String(); id != "" {
				users = append(users, id)
			}
		}
		return users, nil
	}
	return nil, fmt.Errorf("ClickUp workspace %s not visible", service.teamID)
}
