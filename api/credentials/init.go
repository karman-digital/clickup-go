package credentials

import "github.com/hashicorp/go-retryablehttp"

func NewClickUpApiKeyCredentials(apiKey string, teamId string) *Credentials {
	var creds Credentials
	client := retryablehttp.NewClient()
	client.Logger = nil
	creds.client = client
	creds.apiKey = apiKey
	creds.teamId = teamId
	return &creds
}
