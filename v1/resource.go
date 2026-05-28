package v1

// Resource is a combination of id and name.
type Resource struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}
