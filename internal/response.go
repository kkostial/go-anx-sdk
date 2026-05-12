package internal

// RequestWrapper is a wrapper for some anexia api calls.
type RequestWrapper[T any] struct {
	State    string   `json:"state"`
	Messages []string `json:"messages"`
	Data     T        `json:"data"`
}
