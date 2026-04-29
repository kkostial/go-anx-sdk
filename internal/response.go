package internal

type RequestWrapper[T any] struct {
	State    string   `json:"state"`
	Messages []string `json:"messages"`
	Data     T        `json:"data"`
}
