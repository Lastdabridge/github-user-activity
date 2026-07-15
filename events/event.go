package events

type GitHubEvent struct {
	Type string `json:"type"`

	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`

	Payload struct {
		Action string `json:"action"`
	} `json:"payload"`
}