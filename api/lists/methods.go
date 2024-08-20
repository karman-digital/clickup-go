package lists

import (
	"encoding/json"
	"io"
	"net/http"

	listmodels "github.com/karman-digital/clickup/models/lists"
)

func (ls *ListService) CreateList(folderId string, body listmodels.ListCreationBody) (listmodels.List, error) {
	var list listmodels.List
	reqBody, err := json.Marshal(body)
	if err != nil {
		return listmodels.List{}, err
	}
	resp, err := ls.SendRequest(http.MethodPost, "/list", reqBody)
	if err != nil {
		return listmodels.List{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return listmodels.List{}, err
	}
	if resp.StatusCode != http.StatusCreated {
		return listmodels.List{}, err
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
	resp, err := ls.SendRequest(http.MethodPost, "/list", reqBody)
	if err != nil {
		return listmodels.List{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return listmodels.List{}, err
	}
	if resp.StatusCode != http.StatusCreated {
		return listmodels.List{}, err
	}
	err = json.Unmarshal(respBody, &list)
	if err != nil {
		return listmodels.List{}, err
	}
	return list, nil
}
