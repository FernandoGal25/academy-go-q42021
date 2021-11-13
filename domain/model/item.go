package model

// Se deja el nombre generico item en lo que se encuentra un correcto nombre para el modelo de negocio
type Item struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
