package clickupapp

import (
	"github.com/karman-digital/clickup/api/credentials"
	"github.com/karman-digital/clickup/api/lists"
	"github.com/karman-digital/clickup/api/timetracking"
)

func InitClickup() *ClickUp {
	return &ClickUp{}
}

func (c *ClickUp) InitClient(credentials *credentials.Credentials) {
	c.Credentials = credentials
	c.ApiClient = NewApiClient(credentials)
	c.Lists = lists.NewListService(credentials)
}

func NewClickUpInstance(creds *credentials.Credentials) *ClickUp {
	return &ClickUp{
		Credentials: creds,
		ApiClient:   NewApiClient(creds),
	}
}

func NewApiClient(creds *credentials.Credentials) ApiClient {
	return ApiClient{
		TimeTracking: timetracking.NewTimeTracking(creds),
	}
}
