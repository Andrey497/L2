package models

type Event struct {
	Id          int    `json:"id"`
	Name        int    `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}
