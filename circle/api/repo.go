package api

type Repo struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Following bool   `json:"following"`
}
