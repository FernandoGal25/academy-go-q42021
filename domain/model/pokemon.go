package model

/*
	Pokemon entity
*/
type Pokemon struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Order          int    `json:"order"`
	BaseExperience int    `json:"base_experience"`
}
