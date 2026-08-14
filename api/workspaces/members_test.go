package workspaces

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
)

type recordingWorkspaceRequester struct {
	method string
	path   string
}

func (requester *recordingWorkspaceRequester) SendRequestWithContext(_ context.Context, method, path string, _ []byte) (*http.Response, error) {
	requester.method = method
	requester.path = path
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"teams":[{"id":"workspace","members":[{"user":{"id":123}},{"user":{"id":"456"}}]}]}`)),
	}, nil
}

func TestGetWorkspaceUsersSelectsConfiguredWorkspace(t *testing.T) {
	requester := &recordingWorkspaceRequester{}
	service := newWorkspaceService(requester, "workspace")

	users, err := service.GetWorkspaceUserIDs(context.Background())
	if err != nil {
		t.Fatalf("GetWorkspaceUserIDs() error = %v", err)
	}
	if requester.method != http.MethodGet || requester.path != "/team" {
		t.Fatalf("request = %s %s", requester.method, requester.path)
	}
	if len(users) != 2 || users[0] != "123" || users[1] != "456" {
		t.Fatalf("users = %#v", users)
	}
}
