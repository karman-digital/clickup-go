package folders

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/karman-digital/clickup/api/shared"
	foldermodels "github.com/karman-digital/clickup/models/folders"
)

func (service *Service) ListFolders(ctx context.Context, spaceID string, archived bool) ([]foldermodels.Folder, error) {
	path := fmt.Sprintf("/space/%s/folder?archived=%s", url.PathEscape(spaceID), strconv.FormatBool(archived))
	response, err := service.requester.SendRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, classifyTransportError(err)
	}
	var folders foldermodels.FoldersResponse
	if err := decodeResponse(response, &folders); err != nil {
		return nil, err
	}
	return folders.Folders, nil
}

func (service *Service) GetFolder(ctx context.Context, folderID string) (foldermodels.Folder, error) {
	path := fmt.Sprintf("/folder/%s", url.PathEscape(folderID))
	response, err := service.requester.SendRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return foldermodels.Folder{}, classifyTransportError(err)
	}
	var folder foldermodels.Folder
	if err := decodeResponse(response, &folder); err != nil {
		return foldermodels.Folder{}, err
	}
	return folder, nil
}

func (service *Service) ListFolderTemplates(ctx context.Context, workspaceID string) ([]foldermodels.FolderTemplate, error) {
	path, err := buildListFolderTemplatesPath(workspaceID)
	if err != nil {
		return nil, err
	}
	response, err := service.requester.SendRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, classifyTransportError(err)
	}
	var templates foldermodels.FolderTemplatesResponse
	if err := decodeResponse(response, &templates); err != nil {
		return nil, err
	}
	return templates.Templates, nil
}

func (service *Service) MoveFolder(ctx context.Context, folderID, destinationSpaceID string) error {
	folderID = strings.TrimSpace(folderID)
	destinationSpaceID = strings.TrimSpace(destinationSpaceID)
	if folderID == "" {
		return fmt.Errorf("folder ID is required")
	}
	if destinationSpaceID == "" {
		return fmt.Errorf("destination Space ID is required")
	}
	body, err := json.Marshal(foldermodels.MoveFolderRequest{SpaceID: destinationSpaceID})
	if err != nil {
		return fmt.Errorf("encode ClickUp move folder request: %w", err)
	}
	path := fmt.Sprintf("/folder/%s/position", url.PathEscape(folderID))
	response, err := service.requester.SendRequestWithContext(ctx, http.MethodPut, path, body)
	if err != nil {
		return classifyTransportError(err)
	}
	return decodeResponse(response, &struct{}{})
}

func (service *Service) CreateFolderFromTemplate(ctx context.Context, spaceID, templateID string, body foldermodels.CreateFromTemplateBody) (foldermodels.Folder, error) {
	path, requestBody, err := buildCreateFromTemplateRequest(spaceID, templateID, body)
	if err != nil {
		return foldermodels.Folder{}, err
	}
	response, err := service.requester.SendRequestWithContext(ctx, http.MethodPost, path, requestBody)
	if err != nil {
		return foldermodels.Folder{}, classifyTransportError(err)
	}
	var folder foldermodels.Folder
	if err := decodeResponse(response, &folder); err != nil {
		return foldermodels.Folder{}, err
	}
	return folder, nil
}

func buildCreateFromTemplateRequest(spaceID, templateID string, body foldermodels.CreateFromTemplateBody) (string, []byte, error) {
	spaceID = strings.TrimSpace(spaceID)
	templateID = strings.TrimSpace(templateID)
	body.Name = strings.TrimSpace(body.Name)
	if spaceID == "" || templateID == "" || body.Name == "" {
		return "", nil, fmt.Errorf("ClickUp Space, Folder template and name are required")
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return "", nil, fmt.Errorf("encode ClickUp Folder template request: %w", err)
	}
	path := fmt.Sprintf("/space/%s/folder_template/%s", url.PathEscape(spaceID), url.PathEscape(templateID))
	return path, requestBody, nil
}

func buildListFolderTemplatesPath(workspaceID string) (string, error) {
	workspaceID = strings.TrimSpace(workspaceID)
	if workspaceID == "" {
		return "", fmt.Errorf("ClickUp Workspace ID is required")
	}
	return fmt.Sprintf("/team/%s/folder_template", url.PathEscape(workspaceID)), nil
}
func decodeResponse(response *http.Response, target any) error {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read ClickUp folder response: %w", err)
	}
	if err := classifyFolderResponseStatus(response.StatusCode); err != nil {
		return err
	}
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("decode ClickUp folder response: %w", err)
	}
	return nil
}

func classifyFolderResponseStatus(status int) error {
	switch {
	case status == http.StatusOK:
		return nil
	case status == http.StatusNotFound:
		return shared.ErrResourceNotFound
	case status == http.StatusUnauthorized || status == http.StatusForbidden:
		return shared.ErrPermissionDenied
	case status == http.StatusBadRequest || status == http.StatusUnprocessableEntity:
		return fmt.Errorf("%w: ClickUp returned status %d", shared.ErrInvalidRequest, status)
	case status == http.StatusConflict:
		return fmt.Errorf("%w: ClickUp returned status %d", shared.ErrConflict, status)
	case status == http.StatusTooManyRequests || status >= http.StatusInternalServerError:
		return fmt.Errorf("%w: ClickUp returned status %d", shared.ErrTransientFailure, status)
	default:
		return fmt.Errorf("ClickUp Folder request returned status %d", status)
	}
}

func classifyTransportError(err error) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}
	return fmt.Errorf("%w: %v", shared.ErrTransientFailure, err)
}
