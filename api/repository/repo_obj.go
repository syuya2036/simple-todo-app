package repository

import (
	"todo/model"
)

type Repository interface {
	GetAll() ([]model.Task, error)
	GetOne(id int) (model.Task, error)
	CreateOne(task model.Task) (model.Task, error)
	UpdateOne(task model.Task) (model.Task, error)
}