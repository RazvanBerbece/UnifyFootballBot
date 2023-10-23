package apiFootballModels

type League struct {
	Id          int
	Name        string
	CountryName string
	Teams       []Team
}
