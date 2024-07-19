package web

type PondRead struct {
	PondId int    `json:"poind_id"`
	Name   string `json:"name"`
	Owner  string `json:"owner"`
}

type PondCreateRequest struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type PondUpdateRequest struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type PondResponse struct {
	PondId int    `json:"poind_id"`
	Name   string `json:"name"`
	Owner  string `json:"owner"`
}
