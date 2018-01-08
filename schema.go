package main

type Version struct {
	Version	string `json:"version,omitempty"`
}

type Car struct {
	ID           string        `json:"id,omitempty"`
	Make         string        `json:"make,omitempty"`
	Model        string        `json:"model,omitempty"`
	Year         int           `json:"year,omitempty"`
	Transmission *Transmission `json:"transmission,omitempty"`
}

type Transmission struct {
	Type  string `json:"type,omitempty"`
	Gears int `json:"gears,omitempty"`
}

var cars []Car
