package models

type Artist struct {
	Id          int
	Name        string
	Url         string
	Description string
}
type Animation struct {
	// Id     int
	Source []byte `db:"source"`
	// Vibes  []string
	Artist string `db:"artist"`
}

type Song struct {
	Id     int
	Source string
	Name   string
	Vibes  []string
	Artist Artist
}
