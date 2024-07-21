package web

import (
	"log"

	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type PondRead struct {
	PondId int    `json:"pond_id"`
	Name   string `json:"name"`
}

type PondCreateRequest struct {
	Name  string `json:"name"`
	FarmId int `json:"farm_id"`
}

type PondUpdateRequest struct {
	PondId int `json:"pond_id"`
	FarmId int `json:"farm_id"`
	Name  string `json:"name"`
}

func (p *PondUpdateRequest) ChangeSettedField(pond *domain.Pond) {
	log.Println("p: ", p)
	log.Println("pond: ", pond)
	if p.Name != "" {
		pond.Name = p.Name
	}
}

type PondResponse struct {
	PondId int    `json:"pond_id"`
	FarmId int `json:"farm_id"`
	Name   string `json:"name"`
}
