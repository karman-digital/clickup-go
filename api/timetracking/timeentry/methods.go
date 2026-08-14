package timeentry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/karman-digital/clickup/api/shared"
	sharedmodels "github.com/karman-digital/clickup/models/shared"
	timetrackingmodels "github.com/karman-digital/clickup/models/timetracking"
)

func (t *TimeEntryService) GetTimeEntries(ctx context.Context, query timetrackingmodels.TimeEntryQuery) (timetrackingmodels.TimeEntriesResponse, error) {
	parameters := url.Values{}
	if query.StartDate != 0 {
		parameters.Set("start_date", strconv.FormatInt(query.StartDate, 10))
	}
	if query.EndDate != 0 {
		parameters.Set("end_date", strconv.FormatInt(query.EndDate, 10))
	}
	if len(query.AssigneeIDs) != 0 {
		parameters.Set("assignee", strings.Join(query.AssigneeIDs, ","))
	}
	if query.SpaceID != "" {
		parameters.Set("space_id", query.SpaceID)
	}
	if query.IncludeLocationNames {
		parameters.Set("include_location_names", "true")
	}
	path := fmt.Sprintf("/team/%s/time_entries", t.reader.GetTeamId())
	if encoded := parameters.Encode(); encoded != "" {
		path += "?" + encoded
	}
	response, err := t.reader.SendTimeTrackingRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return timetrackingmodels.TimeEntriesResponse{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return timetrackingmodels.TimeEntriesResponse{}, err
	}
	if response.StatusCode != http.StatusOK {
		return timetrackingmodels.TimeEntriesResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	var entries timetrackingmodels.TimeEntriesResponse
	if err := json.Unmarshal(body, &entries); err != nil {
		return timetrackingmodels.TimeEntriesResponse{}, err
	}
	return entries, nil
}

func (t *TimeEntryService) GetTimeEntry(ctx context.Context, id string) (timetrackingmodels.TimeEntryRecordResponse, error) {
	path := fmt.Sprintf("/team/%s/time_entries/%s?include_location_names=true", t.reader.GetTeamId(), url.PathEscape(id))
	response, err := t.reader.SendTimeTrackingRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return timetrackingmodels.TimeEntryRecordResponse{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return timetrackingmodels.TimeEntryRecordResponse{}, err
	}
	if response.StatusCode == http.StatusNotFound {
		return timetrackingmodels.TimeEntryRecordResponse{}, shared.ErrResourceNotFound
	}
	if response.StatusCode != http.StatusOK {
		return timetrackingmodels.TimeEntryRecordResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	var entry timetrackingmodels.TimeEntryRecordResponse
	if err := json.Unmarshal(body, &entry); err != nil {
		return timetrackingmodels.TimeEntryRecordResponse{}, err
	}
	return entry, nil
}

func (t *TimeEntryService) GetTimeEntryHistory(id string) (timetrackingmodels.TimeTrackHistoryResponse, error) {
	var timeTrackHistory timetrackingmodels.TimeTrackHistoryResponse
	resp, err := t.SendTimeTrackingRequest(http.MethodGet, fmt.Sprintf("/team/%s/time_entries/%s/history", t.GetTeamId(), id), nil)
	if err != nil {
		return timetrackingmodels.TimeTrackHistoryResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return timetrackingmodels.TimeTrackHistoryResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return timetrackingmodels.TimeTrackHistoryResponse{}, shared.ErrResourceNotFound
		}
		return timetrackingmodels.TimeTrackHistoryResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	err = json.Unmarshal(body, &timeTrackHistory)
	if err != nil {
		return timetrackingmodels.TimeTrackHistoryResponse{}, err
	}
	return timeTrackHistory, nil
}

func (t *TimeEntryService) CreateTimeEntry(timeEntry timetrackingmodels.TimeEntry, opts ...sharedmodels.GetOptions) (timetrackingmodels.TimeEntryResponse, error) {
	responseTimeEntry := timetrackingmodels.TimeEntryResponse{}
	requestBody, err := json.Marshal(timeEntry)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	resp, err := t.SendTimeTrackingRequest(http.MethodPost, fmt.Sprintf("/team/%s/time_entries", t.GetTeamId()), requestBody, opts...)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return timetrackingmodels.TimeEntryResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	err = json.Unmarshal(body, &responseTimeEntry)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	return responseTimeEntry, nil
}

func (t *TimeEntryService) DeleteTimeEntry(id string) error {
	resp, err := t.SendTimeTrackingRequest(http.MethodDelete, fmt.Sprintf("/team/%s/time_entries/%s", t.GetTeamId(), id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error body: %s", string(body))
	}
	return nil
}
