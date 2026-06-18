package models

type Student struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Marks int `json:"marks"`
	Status string `json:"status"`
}