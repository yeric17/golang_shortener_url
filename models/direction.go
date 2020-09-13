package models

import (
	"encoding/json"
	"io"
	"time"
	//_ "url_shortener/controllers"
)

// Direction es la entidad de una direccion a redirigir
type Direction struct {
	ID       uint64    `json:"id"`
	URL      string    `json:"url"`
	ShortURL string    `json:"short_url"`
	CreateAt time.Time `-`
	UpdateAt time.Time `-`
	DeleteAt time.Time `-`
}

// ToJSON devuelve una respuesta codificada en json
func (d *Direction) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(d)
}

// FromJSON obtiene del cliente una respuesta codificada en json
func (d *Direction) FromJSON(r io.Reader) error {
	encoder := json.NewDecoder(r)
	return encoder.Decode(d)
}
