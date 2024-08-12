package taskdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/karman-digital/clickup/api/shared"
	taskmodels "github.com/karman-digital/clickup/models/tasks"
)

func (t *TaskDataService) GetTask(taskID string, opts ...taskmodels.GetTaskOptions) (taskmodels.Task, error) {
	resp, err := t.SendTaskRequest(http.MethodGet, fmt.Sprintf("/task/%s", taskID), nil, opts...)
	if err != nil {
		return taskmodels.Task{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return taskmodels.Task{}, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return taskmodels.Task{}, shared.ErrResourceNotFound
		}
		return taskmodels.Task{}, fmt.Errorf("error body: %s", string(body))
	}
	var task taskmodels.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		return taskmodels.Task{}, err
	}
	return task, nil
}
