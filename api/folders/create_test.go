package folders

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/karman-digital/clickup/api/shared"
	foldermodels "github.com/karman-digital/clickup/models/folders"
)

func TestBuildCreateFromTemplateRequest(t *testing.T) {
	path, body, err := buildCreateFromTemplateRequest("space/1", "template/1", foldermodels.CreateFromTemplateBody{Name: "  Acme Ltd  "})
	if err != nil {
		t.Fatalf("buildCreateFromTemplateRequest() error = %v", err)
	}
	if path != "/space/space%2F1/folder_template/template%2F1" {
		t.Fatalf("path = %q", path)
	}
	var decoded foldermodels.CreateFromTemplateBody
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if decoded.Name != "Acme Ltd" {
		t.Fatalf("name = %q", decoded.Name)
	}
}

func TestBuildCreateFromTemplateRequestRejectsMissingValues(t *testing.T) {
	for _, test := range []struct {
		name       string
		spaceID    string
		templateID string
		folderName string
	}{
		{name: "space", templateID: "template", folderName: "Company"},
		{name: "template", spaceID: "space", folderName: "Company"},
		{name: "folder name", spaceID: "space", templateID: "template"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, _, err := buildCreateFromTemplateRequest(test.spaceID, test.templateID, foldermodels.CreateFromTemplateBody{Name: test.folderName}); err == nil {
				t.Fatal("buildCreateFromTemplateRequest() error = nil")
			}
		})
	}
}

func TestBuildListFolderTemplatesPath(t *testing.T) {
	path, err := buildListFolderTemplatesPath(" workspace/1 ")
	if err != nil {
		t.Fatalf("buildListFolderTemplatesPath() error = %v", err)
	}
	if path != "/team/workspace%2F1/folder_template" {
		t.Fatalf("path = %q", path)
	}
	if _, err := buildListFolderTemplatesPath(" "); err == nil {
		t.Fatal("buildListFolderTemplatesPath() error = nil")
	}
}

func TestClassifyFolderResponseStatus(t *testing.T) {
	for _, test := range []struct {
		status int
		want   error
	}{
		{status: http.StatusNotFound, want: shared.ErrResourceNotFound},
		{status: http.StatusUnauthorized, want: shared.ErrPermissionDenied},
		{status: http.StatusForbidden, want: shared.ErrPermissionDenied},
		{status: http.StatusBadRequest, want: shared.ErrInvalidRequest},
		{status: http.StatusUnprocessableEntity, want: shared.ErrInvalidRequest},
		{status: http.StatusConflict, want: shared.ErrConflict},
		{status: http.StatusTooManyRequests, want: shared.ErrTransientFailure},
		{status: http.StatusInternalServerError, want: shared.ErrTransientFailure},
	} {
		t.Run(http.StatusText(test.status), func(t *testing.T) {
			if err := classifyFolderResponseStatus(test.status); !errors.Is(err, test.want) {
				t.Fatalf("classifyFolderResponseStatus(%d) = %v, want %v", test.status, err, test.want)
			}
		})
	}
	if err := classifyFolderResponseStatus(http.StatusOK); err != nil {
		t.Fatalf("classifyFolderResponseStatus(200) = %v", err)
	}
}
