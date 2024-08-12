package credentials

import "github.com/hashicorp/go-retryablehttp"

type Credentials struct {
	client *retryablehttp.Client
	apiKey string
	teamId string
}
