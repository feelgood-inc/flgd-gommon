package models

type HealthCareFacility struct {
	ID        uint    `json:"id"`
	UID       string  `json:"uid"`
	Name      string  `json:"name"`
	Street    string  `json:"address"`
	Number    string  `json:"number"`
	Commune   string  `json:"commune"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
