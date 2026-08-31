package lists

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetListAndFolderlessListsUseReadEndpoints(t *testing.T) {
	paths := []string{}
	service := &ListService{send: func(method, path string, _ []byte) (*http.Response, error) {
		if method != http.MethodGet {
			t.Fatalf("method = %s", method)
		}
		paths = append(paths, path)
		body := `{"id":"123","name":"PJT03000 Job"}`
		if strings.Contains(path, "/space/") {
			body = `{"lists":[{"id":"123","name":"PJT03000 Job"}]}`
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body))}, nil
	}}
	list, err := service.GetList("123")
	if err != nil || list.ID != "123" {
		t.Fatalf("GetList() = %#v, %v", list, err)
	}
	lists, err := service.GetFolderlessLists("space", false)
	if err != nil || len(lists) != 1 || lists[0].ID != "123" {
		t.Fatalf("GetFolderlessLists() = %#v, %v", lists, err)
	}
	if len(paths) != 2 || paths[0] != "/list/123" || paths[1] != "/space/space/list?archived=false" {
		t.Fatalf("paths = %#v", paths)
	}
}
