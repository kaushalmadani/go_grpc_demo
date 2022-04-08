package models
type Filter struct {
	Key string `json:"key"`

	// If value is empty, just return everything but sorted with the key
	Value string `json:"value,omitempty"`
}

type Status string

const (
	Success Status = "Success"
	Failed  Status = "Failed"
)