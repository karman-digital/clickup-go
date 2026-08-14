package timeentry

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/karman-digital/clickup/api/shared"
	timetrackingmodels "github.com/karman-digital/clickup/models/timetracking"
)

type recordingTimeEntryRequester struct {
	status int
	body   string
	method string
	path   string
}

func (requester *recordingTimeEntryRequester) GetTeamId() string {
	return "workspace"
}

func (requester *recordingTimeEntryRequester) SendTimeTrackingRequestWithContext(_ context.Context, method, path string, _ []byte) (*http.Response, error) {
	requester.method = method
	requester.path = path
	return &http.Response{
		StatusCode: requester.status,
		Body:       io.NopCloser(bytes.NewBufferString(requester.body)),
	}, nil
}

func TestGetTimeEntriesBuildsBoundedWorkspaceQueryAndDecodesLocations(t *testing.T) {
	requester := &recordingTimeEntryRequester{
		status: http.StatusOK,
		body:   `{"data":[{"id":"entry-1","user":{"id":123},"start":"1722510000000","end":"1722513600000","duration":"3600000","task_location":{"list_id":456,"space_id":90152724343}}]}`,
	}
	service := newTimeEntryService(requester)

	entries, err := service.GetTimeEntries(context.Background(), timetrackingmodels.TimeEntryQuery{
		StartDate:            1722510000000,
		EndDate:              1722513600000,
		AssigneeIDs:          []string{"123", "456"},
		SpaceID:              "90152724343",
		IncludeLocationNames: true,
	})
	if err != nil {
		t.Fatalf("GetTimeEntries() error = %v", err)
	}
	if requester.method != http.MethodGet {
		t.Fatalf("method = %q, want GET", requester.method)
	}
	wantPath := "/team/workspace/time_entries?assignee=123%2C456&end_date=1722513600000&include_location_names=true&space_id=90152724343&start_date=1722510000000"
	if requester.path != wantPath {
		t.Fatalf("path = %q, want %q", requester.path, wantPath)
	}
	if len(entries.Data) != 1 {
		t.Fatalf("len(entries) = %d, want 1", len(entries.Data))
	}
	entry := entries.Data[0]
	if entry.ID.String() != "entry-1" || entry.User.ID.String() != "123" || entry.TaskLocation.ListID.String() != "456" || entry.TaskLocation.SpaceID.String() != "90152724343" {
		t.Fatalf("entry = %#v", entry)
	}
}

func TestGetTimeEntryReturnsResourceNotFound(t *testing.T) {
	requester := &recordingTimeEntryRequester{status: http.StatusNotFound, body: `{"err":"not found"}`}
	service := newTimeEntryService(requester)

	_, err := service.GetTimeEntry(context.Background(), "missing")
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("GetTimeEntry() error = %v, want ErrResourceNotFound", err)
	}
	if requester.path != "/team/workspace/time_entries/missing?include_location_names=true" {
		t.Fatalf("path = %q", requester.path)
	}
}
