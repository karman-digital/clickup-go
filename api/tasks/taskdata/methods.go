package taskdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/karman-digital/clickup/api/shared"
	taskmodels "github.com/karman-digital/clickup/models/tasks"
)

func (t *TaskDataService) GetTask(taskID string, opts ...taskmodels.GetTaskOptions) (taskmodels.TaskGetResponse, error) {
	resp, err := t.SendTaskRequest(http.MethodGet, fmt.Sprintf("/task/%s", taskID), nil, opts...)
	if err != nil {
		return taskmodels.TaskGetResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return taskmodels.TaskGetResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return taskmodels.TaskGetResponse{}, shared.ErrResourceNotFound
		}
		return taskmodels.TaskGetResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	var task taskmodels.TaskGetResponse
	err = json.Unmarshal(body, &task)
	if err != nil {
		return taskmodels.TaskGetResponse{}, err
	}
	return task, nil
}

func (t *TaskDataService) CreateTask(listId string, task taskmodels.TaskPostBody) (taskmodels.TaskGetResponse, error) {
	postBody, err := json.Marshal(task)
	if err != nil {
		return taskmodels.TaskGetResponse{}, err
	}
	resp, err := t.SendTaskRequest(http.MethodPost, fmt.Sprintf("/list/%s/task", listId), postBody)
	if err != nil {
		return taskmodels.TaskGetResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return taskmodels.TaskGetResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return taskmodels.TaskGetResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	var taskResp taskmodels.TaskGetResponse
	err = json.Unmarshal(body, &taskResp)
	if err != nil {
		return taskmodels.TaskGetResponse{}, err
	}
	return taskResp, nil
}
