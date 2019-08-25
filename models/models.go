package models

type Artist struct {
	Id          int
	Name        string
	Url         string
	Description string
}

type Animation struct {
	Id     int
	Source string
	Vibes  []string
	Artist Artist
}

type Song struct {
	Id     int
	Source string
	Name   string
	Vibes  []string
	Artist Artist
}
