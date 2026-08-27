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

func decodeResponse(response *http.Response, target any) error {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read ClickUp folder response: %w", err)
	}
	switch {
	case response.StatusCode == http.StatusNotFound:
		return shared.ErrResourceNotFound
	case response.StatusCode == http.StatusUnauthorized || response.StatusCode == http.StatusForbidden:
		return shared.ErrPermissionDenied
	case response.StatusCode == http.StatusTooManyRequests || response.StatusCode >= http.StatusInternalServerError:
		return fmt.Errorf("%w: ClickUp returned status %d", shared.ErrTransientFailure, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("ClickUp folder request returned status %d: %s", response.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("decode ClickUp folder response: %w", err)
	}
	return nil
}

func classifyTransportError(err error) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}
	return fmt.Errorf("%w: %v", shared.ErrTransientFailure, err)
}
