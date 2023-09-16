package repository

import (
	"todo/model"
)

type repository struct {
	SqlHandler // gorm.DB
}

func NewRepository(sqlHandler SqlHandler) Repository {
	repository := repository{sqlHandler}
	return &repository
}

func (repo *repository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	if err := repo.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *repository) GetOne(id int) (model.Task, error) {
	var task model.Task
	if err := repo.First(&task, id).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (repo *repository) CreateOne(task model.Task) (model.Task, error) {
	if err := repo.Create(&task).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}


func (repo *repository) UpdateOne(task model.Task) (model.Task, error) {
	if err := repo.Save(&task).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}