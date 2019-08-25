package models

type Artist struct {
	Id          int
	name        string
	url         string
	description string
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
