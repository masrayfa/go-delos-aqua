package web

type FarmRead struct {
	FarmId   int    `json:"farm_id"`
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type FarmRequest struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type FarmResponse struct {
	FarmId   int    `json:"farm_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}