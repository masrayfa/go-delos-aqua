package web

import "github.com/masrayfa/go-delos-aqua/internals/models/domain"

type PondRead struct {
	PondId int    `json:"poind_id"`
	Name   string `json:"name"`
}

type PondCreateRequest struct {
	Name  string `json:"name"`
}

type PondUpdateRequest struct {
	PondId int `json:"pond_id"`
	Name  string `json:"name"`
}

func (p *PondUpdateRequest) ChangeSettedField(pond *domain.Pond) {
	if p.Name != "" {
		pond.Name = p.Name
	}
}

type PondResponse struct {
	PondId int    `json:"poind_id"`
	Name   string `json:"name"`
}
