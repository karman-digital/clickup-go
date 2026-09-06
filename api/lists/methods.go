package lists

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/karman-digital/clickup/api/shared"
	listmodels "github.com/karman-digital/clickup/models/lists"
)

func (ls *ListService) CreateList(folderId string, body listmodels.ListCreationBody) (listmodels.List, error) {
	var list listmodels.List
	reqBody, err := json.Marshal(body)
	if err != nil {
		return listmodels.List{}, err
	}
	resp, err := ls.sendRequest(http.MethodPost, fmt.Sprintf("/folder/%s/list", folderId), reqBody)
	if err != nil {
		return listmodels.List{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return listmodels.List{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return listmodels.List{}, errors.New("error creating list " + string(respBody))
	}
	err = json.Unmarshal(respBody, &list)
	if err != nil {
		return listmodels.List{}, err
	}
	return list, nil
}

func (ls *ListService) CreateFolderlessList(spaceId string, body listmodels.ListCreationBody) (listmodels.List, error) {
	var list listmodels.List
	reqBody, err := json.Marshal(body)
	if err != nil {
		return listmodels.List{}, err
	}
	resp, err := ls.sendRequest(http.MethodPost, fmt.Sprintf("/space/%s/list", spaceId), reqBody)
	if err != nil {
		return listmodels.List{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return listmodels.List{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return listmodels.List{}, errors.New("error creating list " + string(respBody))
	}
	fmt.Println(string(respBody))
	err = json.Unmarshal(respBody, &list)
	if err != nil {
		return listmodels.List{}, err
	}
	return list, nil
}

func (ls *ListService) CreateFolderlessListFromTemplate(spaceId string, templateId string, body listmodels.ListCreationBody) (listmodels.List, error) {
	var list listmodels.List
	reqBody, err := json.Marshal(body)
	if err != nil {
		return listmodels.List{}, err
	}
	resp, err := ls.sendRequest(http.MethodPost, fmt.Sprintf("/space/%s/list_template/%s", spaceId, templateId), reqBody)
	if err != nil {
		return listmodels.List{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return listmodels.List{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return listmodels.List{}, errors.New("error creating list from template " + string(respBody))
	}
	err = json.Unmarshal(respBody, &list)
	if err != nil {
		return listmodels.List{}, err
	}
	return list, nil
}

func (ls *ListService) GetList(listID string) (listmodels.List, error) {
	var list listmodels.List
	resp, err := ls.sendRequest(http.MethodGet, fmt.Sprintf("/list/%s", url.PathEscape(listID)), nil)
	if err != nil {
		return list, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return list, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return list, shared.ErrResourceNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return list, fmt.Errorf("error getting list: status %d: %s", resp.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, &list); err != nil {
		return list, err
	}
	return list, nil
}

func (ls *ListService) GetFolderlessLists(spaceID string, archived bool) ([]listmodels.List, error) {
	path := fmt.Sprintf("/space/%s/list?archived=%t", url.PathEscape(spaceID), archived)
	resp, err := ls.sendRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting folderless lists: status %d: %s", resp.StatusCode, string(body))
	}
	var result listmodels.ListsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Lists, nil
}

func (ls *ListService) sendRequest(method, path string, body []byte) (*http.Response, error) {
	if ls.send != nil {
		return ls.send(method, path, body)
	}
	return ls.SendRequest(method, path, body)
}
