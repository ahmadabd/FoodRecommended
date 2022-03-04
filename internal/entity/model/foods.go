package model

import "github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"

type Food struct {
	Id         int             `json:"id"`
	Name       string          `json:"name"`
	City       string          `json:"city"`
	Country    string          `json:"country"`
	Vegetarian enum.Vegetative `json:"vegetarian"`
}

// CREATE TABLE foods ( id int auto_increment, name nvarchar(20) not null, country nvarchar(20), city nvarchar(20), primary key(id));
