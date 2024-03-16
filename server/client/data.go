package client

import "example/company/server/branch"

type Client struct {
	ClientID   int64         `json:"client_id"`
	ClientName string        `json:"client_name"`
	Branch     branch.Branch `json:"branch"`
}
