package alerts

import (
	"os"
)

type AcknowledgeAlertRequest struct {
	ApiKey string `json:"apiKey,omitempty"`
	Id     string `json:"id,omitempty"`
	Alias  string `json:"alias,omitempty"`
	User   string `json:"user,omitempty"`
	Note   string `json:"note,omitempty"`
	Source string `json:"source,omitempty"`
}

type AddNoteAlertRequest struct {
	ApiKey string `json:"apiKey,omitempty"`
	Id     string `json:"id,omitempty"`
	Alias  string `json:"alias,omitempty"`
	Note   string `json:"note,omitempty"`
	User   string `json:"user,omitempty"`
	Source string `json:"source,omitempty"`
}

type AddRecipientAlertRequest struct {
	ApiKey    string `json:"apiKey,omitempty"`
	Id        string `json:"id,omitempty"`
	Alias     string `json:"alias,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	User      string `json:"user,omitempty"`
	Note      string `json:"note,omitempty"`
	Source    string `json:"source,omitempty"`
}

type AddTeamAlertRequest struct {
	ApiKey string `json:"apiKey,omitempty"`
	Id     string `json:"id,omitempty"`
	Alias  string `json:"alias,omitempty"`
	Team   string `json:"team,omitempty"`
	User   string `json:"user,omitempty"`
	Note   string `json:"note,omitempty"`
	Source string `json:"source,omitempty"`
}
type AssignOwnerAlertRequest struct {
	ApiKey string `json:"apiKey,omitempty"`
	Id     string `json:"id,omitempty"`
	Alias  string `json:"alias,omitempty"`
	Owner  string `json:"owner,omitempty"`
	User   string `json:"user,omitempty"`
	Note   string `json:"note,omitempty"`
	Source string `json:"source,omitempty"`
}

type AttachFileAlertRequest struct {
	ApiKey     string   `json:"apiKey,omitempty"`
	Id         string   `json:"id,omitempty"`
	Alias      string   `json:"alias,omitempty"`
	Attachment *os.File `json:"attachment,omitempty"`
	User       string   `json:"user,omitempty"`
	Source     string   `json:"source,omitempty"`
	IndexFile  string   `json:"indexFile,omitempty"`
	Note       string   `json:"note,omitempty"`
}

type CloseAlertRequest struct {
	ApiKey string   `json:"apiKey,omitempty"`
	Id     string   `json:"id,omitempty"`
	Alias  string   `json:"alias,omitempty"`
	User   string   `json:"user,omitempty"`
	Note   string   `json:"note,omitempty"`
	Notify []string `json:"notify,omitempty"`
	Source string   `json:"source,omitempty"`
}

type CreateAlertRequest struct {
	ApiKey      string            `json:"apiKey,omitempty"`
	Message     string            `json:"message,omitempty"`
	Teams       []string          `json:"teams,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Description string            `json:"description,omitempty"`
	Recipients  []string          `json:"recipients,omitempty"`
	Actions     []string          `json:"actions,omitempty"`
	Source      string            `json:"source,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Details     map[string]string `json:"details,omitempty"`
	Entity      string            `json:"entity,omitempty"`
	User        string            `json:"user,omitempty"`
	Note        string            `json:"note,omitempty"`
}

type DeleteAlertRequest struct {
	ApiKey string `url:"apiKey,omitempty"`
	Id     string `url:"id,omitempty"`
	Alias  string `url:"alias,omitempty"`
	User   string `url:"user,omitempty"`
	Source string `url:"source,omitempty"`
}

type ExecuteActionAlertRequest struct {
	ApiKey string `json:"apiKey,omitempty"`
	Id     string `json:"id,omitempty"`
	Alias  string `json:"alias,omitempty"`
	Action string `json:"action,omitempty"`
	User   string `json:"user,omitempty"`
	Source string `json:"source,omitempty"`
	Note   string `json:"note,omitempty"`
}

type GetAlertRequest struct {
	ApiKey string `url:"apiKey,omitempty"`
	Id     string `url:"id,omitempty"`
	Alias  string `url:"alias,omitempty"`
	TinyId string `url:"tinyId,omitempty"`
}

type ListAlertLogsRequest struct {
	ApiKey  string `url:"apiKey,omitempty"`
	Id      string `url:"id,omitempty"`
	Alias   string `url:"alias,omitempty"`
	Limit   string `url:"limit,omitempty"`
	Order   string `url:"order,omitempty"`
	LastKey string `url:"lastKey,omitempty"`
}

type ListAlertNotesRequest struct {
	ApiKey  string `url:"apiKey,omitempty"`
	Id      string `url:"id,omitempty"`
	Alias   string `url:"alias,omitempty"`
	Limit   string `url:"limit,omitempty"`
	Order   string `url:"order,omitempty"`
	LastKey string `url:"lastKey,omitempty"`
}

type ListAlertRecipientsRequest struct {
	ApiKey string `url:"apiKey,omitempty"`
	Id     string `url:"id,omitempty"`
	Alias  string `url:"alias,omitempty"`
}

type ListAlertsRequest struct {
	ApiKey        string `url:"apiKey,omitempty"`
	CreatedAfter  uint64 `url:"createdAfter,omitempty"`
	CreatedBefore uint64 `url:"createdBefore,omitempty"`
	UpdatedAfter  uint64 `url:"updatedAfter,omitempty"`
	UpdatedBefore uint64 `url:"updatedBefore,omitempty"`
	Limit         uint64 `url:"limit,omitempty"`
	Status        string `url:"status,omitempty"`
	SortBy        string `url:"sortBy,omitempty"`
	Order         string `url:"order,omitempty"`
}

type RenotifyAlertRequest struct {
	ApiKey     string   `json:"apiKey,omitempty"`
	Id         string   `json:"id,omitempty"`
	Alias      string   `json:"alias,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	User       string   `json:"user,omitempty"`
	Note       string   `json:"note,omitempty"`
	Source     string   `json:"source,omitempty"`
}

type TakeOwnershipAlertRequest struct {
	ApiKey string `json:"apiKey,omitempty"`
	Id     string `json:"id,omitempty"`
	Alias  string `json:"alias,omitempty"`
	User   string `json:"user,omitempty"`
	Note   string `json:"note,omitempty"`
	Source string `json:"source,omitempty"`
}
